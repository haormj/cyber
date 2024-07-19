package service

import (
	"errors"
	"sync"
	"time"

	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/pb"
	"github.com/haormj/cyber/transport"
	"github.com/haormj/cyber/transport/identity"
	"github.com/haormj/cyber/transport/message"
	"github.com/haormj/cyber/transport/qos"
	"github.com/haormj/cyber/transport/receiver"
	"github.com/haormj/cyber/transport/transmitter"
	"google.golang.org/protobuf/proto"
)

type ResposneCallback[Response proto.Message] func(response Response)

type Client[Request, Response proto.Message] struct {
	nodeName             string
	serviceName          string
	requestChannel       string
	responseChannel      string
	requestTransmitter   transmitter.Transmitter[Request]
	responseReceiver     receiver.Receiver[Response]
	writerID             *identity.Identity
	sequenceNumber       uint64
	pendingRequestsMutex sync.Mutex
	pendingRequests      map[uint64]ResposneCallback[Response]
}

func NewClient[Request, Response proto.Message](nodeName, serviceName string) *Client[Request, Response] {
	return &Client[Request, Response]{
		nodeName:        nodeName,
		serviceName:     serviceName,
		requestChannel:  serviceName + common.SRV_CHANNEL_REQ_SUFFIX,
		responseChannel: serviceName + common.SRV_CHANNEL_RES_SUFFIX,
		pendingRequests: make(map[uint64]ResposneCallback[Response]),
	}
}

func (c *Client[Request, Response]) handleResponse(response Response, msgInfo *message.MessageInfo) {
	log.Logger.Debug("client receive response", "channel", c.responseChannel)
	c.pendingRequestsMutex.Lock()
	defer c.pendingRequestsMutex.Unlock()

	if identity.NotEqual(c.writerID, msgInfo.SpareID) {
		return
	}

	v, ok := c.pendingRequests[msgInfo.SeqNum]
	if !ok {
		return
	}

	v(response)
	delete(c.pendingRequests, msgInfo.SeqNum)
}

func (c *Client[Request, Response]) isInit() bool {
	return c.responseReceiver != nil
}

func (c *Client[Request, Response]) ServiceName() string {
	return c.serviceName
}

func (c *Client[Request, Response]) Init() bool {
	roleAttr := &pb.RoleAttributes{}
	roleAttr.HostName = proto.String(common.GlobalDataInstance.HostName())
	roleAttr.HostIp = proto.String(common.GlobalDataInstance.HostIP())
	roleAttr.ProcessId = proto.Int32(int32(common.GlobalDataInstance.ProcessID()))
	roleAttr.NodeName = proto.String(c.nodeName)
	roleAttr.ChannelName = proto.String(c.requestChannel)
	requestChannelID := common.GlobalDataInstance.RegisterChannel(c.requestChannel)
	roleAttr.ChannelId = proto.Uint64(requestChannelID)
	qosProfile := qos.QOS_PROFILE_SERVICES_DEFAULT
	roleAttr.QosProfile = &qosProfile
	requestTransmitter, err := transport.CreateTransmitter[Request](roleAttr, pb.OptionalMode_SHM)
	if err != nil {
		log.Logger.Error("create reqeust transmitter error", "err", err)
		return false
	}
	c.requestTransmitter = requestTransmitter
	c.writerID = identity.CloneIdentity(requestTransmitter.ID())

	roleAttr.ChannelName = proto.String(c.responseChannel)
	responseChannelID := common.GlobalDataInstance.RegisterChannel(c.responseChannel)
	roleAttr.ChannelId = proto.Uint64(responseChannelID)
	responseReceiver, err := transport.CreateReceiver[Response](roleAttr, pb.OptionalMode_SHM,
		func(msg Response, msgInfo *message.MessageInfo, attr *pb.RoleAttributes) {
			c.handleResponse(msg, msgInfo)
		})
	if err != nil {
		log.Logger.Error("create response receiver error", "err", err)
		return false
	}
	c.responseReceiver = responseReceiver

	return true
}

func (c *Client[Request, Response]) SendRequest(request Request, timeout time.Duration) (Response, error) {
	response := common.Zero[Response]()
	if !c.isInit() {
		return response, errors.New("client not init")
	}
	ch := make(chan Response)
	defer close(ch)

	if err := c.AsyncSendRequest(request, func(r Response) {
		select {
		case ch <- r:
		default:
		}
	}); err != nil {
		return response, err
	}

	select {
	case response := <-ch:
		return response, nil
	case <-time.After(timeout):
		return response, errors.New("request timeout")
	}
}

func (c *Client[Request, Response]) AsyncSendRequest(request Request, ResposneCallback ResposneCallback[Response]) error {
	if !c.isInit() {
		return errors.New("client not init")
	}

	c.pendingRequestsMutex.Lock()
	defer c.pendingRequestsMutex.Unlock()
	c.sequenceNumber++
	msgInfo := message.NewMessageInfo()
	msgInfo.SeqNum = c.sequenceNumber
	msgInfo.SenderID = c.writerID
	msgInfo.SpareID = c.writerID

	if err := c.requestTransmitter.TransmitWithMessageInfo(request, msgInfo); err != nil {
		return err
	}

	c.pendingRequests[c.sequenceNumber] = ResposneCallback

	return nil
}
