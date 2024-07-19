package service

import (
	"sync"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/state"
	"github.com/haormj/cyber/transport"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/qos"
	"github.com/haormj/cyber/transport/receiver"
	"github.com/haormj/cyber/transport/transmitter"
	"google.golang.org/protobuf/proto"
)

type ServiceCallback[Request, Response proto.Message] func(request Request, response *Response)

type Service[Request, Response proto.Message] struct {
	inited                    bool
	nodeName                  string
	serviceName               string
	requestChannel            string
	responseChannel           string
	serviceCallback           ServiceCallback[Request, Response]
	requestReceiver           receiver.Receiver[Request]
	responseTransmitter       transmitter.Transmitter[Response]
	serviceHandleRequestMutex sync.Mutex
	queueMutex                sync.Mutex
	queueCond                 *sync.Cond
	tasks                     []func()
	wg                        sync.WaitGroup
}

func NewService[Request, Response proto.Message](nodeName, serviceName string,
	serviceCallback ServiceCallback[Request, Response]) *Service[Request, Response] {
	s := &Service[Request, Response]{
		nodeName:        nodeName,
		serviceName:     serviceName,
		requestChannel:  serviceName + common.SRV_CHANNEL_REQ_SUFFIX,
		responseChannel: serviceName + common.SRV_CHANNEL_RES_SUFFIX,
		serviceCallback: serviceCallback,
	}
	s.queueCond = sync.NewCond(&s.queueMutex)
	return s
}

func (s *Service[Request, Response]) ServiceName() string {
	return s.serviceName
}

func (s *Service[Request, Response]) Init() bool {
	if s.isInit() {
		return true
	}

	roleAttr := &pb.RoleAttributes{}
	roleAttr.HostName = proto.String(common.GlobalDataInstance.HostName())
	roleAttr.HostIp = proto.String(common.GlobalDataInstance.HostIP())
	roleAttr.ProcessId = proto.Int32(int32(common.GlobalDataInstance.ProcessID()))
	roleAttr.NodeName = proto.String(s.nodeName)
	roleAttr.ChannelName = proto.String(s.responseChannel)
	responseChannelID := common.GlobalDataInstance.RegisterChannel(s.responseChannel)
	roleAttr.ChannelId = proto.Uint64(responseChannelID)
	qosProfile := qos.QOS_PROFILE_SERVICES_DEFAULT
	roleAttr.QosProfile = &qosProfile

	responseTransmitter, err := transport.CreateTransmitter[Response](roleAttr, pb.OptionalMode_SHM)
	if err != nil {
		log.Logger.Error("create response transmitter error", "err", err)
		return false
	}
	s.responseTransmitter = responseTransmitter

	roleAttr.ChannelName = proto.String(s.requestChannel)
	requestChannelID := common.GlobalDataInstance.RegisterChannel(s.requestChannel)
	roleAttr.ChannelId = proto.Uint64(requestChannelID)
	requestReceiver, err := transport.CreateReceiver[Request](roleAttr, pb.OptionalMode_SHM,
		func(request Request, msgInfo *message.MessageInfo, attr *pb.RoleAttributes) {
			s.enqueue(func() {
				s.handleRequest(request, msgInfo)
			})
		})
	if err != nil {
		log.Logger.Error("request receiver craete failed", "err", err)
		return false
	}
	s.requestReceiver = requestReceiver
	s.inited = true
	s.wg.Add(1)
	go s.process()
	return true
}

func (s *Service[Request, Response]) Destroy() {
	s.inited = false
	s.queueCond.Broadcast()
	s.wg.Wait()
}

func (s *Service[Request, Response]) handleRequest(request Request, msgInfo *message.MessageInfo) {
	if !s.isInit() {
		return
	}

	log.Logger.Debug("handing request", "requestChannel", s.requestChannel)
	s.serviceHandleRequestMutex.Lock()
	defer s.serviceHandleRequestMutex.Unlock()

	response := common.Zero[*Response]()
	s.serviceCallback(request, response)
	responseMsgInfo := message.CloneMessageInfo(msgInfo)
	responseMsgInfo.SenderID = s.responseTransmitter.ID()
	s.sendResponse(responseMsgInfo, *response)
}

func (s *Service[Request, Response]) sendResponse(msgInfo *message.MessageInfo, response Response) {
	if !s.isInit() {
		return
	}

	s.responseTransmitter.TransmitWithMessageInfo(response, msgInfo)
}

func (s *Service[Request, Response]) isInit() bool {
	return s.requestReceiver != nil
}

func (s *Service[Request, Response]) enqueue(task func()) {
	s.queueMutex.Lock()
	defer s.queueMutex.Unlock()

	s.tasks = append(s.tasks, task)
	s.queueCond.Signal()
}

func (s *Service[Request, Response]) process() {
	defer s.wg.Done()

	var task func()
	for !state.IsShutdown() {
		s.queueMutex.Lock()
		if !s.inited || len(s.tasks) == 0 {
			s.queueCond.Wait()
		}

		if !s.inited {
			log.Logger.Error("service not init, exit")
			break
		}

		if len(s.tasks) != 0 {
			task = s.tasks[0]
			s.tasks = s.tasks[1:]
		}
		s.queueMutex.Unlock()

		task()
	}
}
