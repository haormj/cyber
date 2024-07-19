package component

import (
	"github.com/haormj/cyber/common"
	"github.com/haormj/cyber/data"
	"github.com/haormj/cyber/log"
	"github.com/haormj/cyber/node"
	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type DataDrivenComponent struct {
	node      node.Node
	component Component
}

func NewDataDrivenComponent(component Component) *DataDrivenComponent {
	return &DataDrivenComponent{
		component: component,
	}
}

func (c *DataDrivenComponent) Initialize(config *pb.ComponentConfig) bool {
	if len(config.GetName()) == 0 {
		log.Logger.Error("name empty")
		return false
	}
	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		log.Logger.Error("component init failed")
		return false
	}

	return true
}

type DataDrivenComponent1[M proto.Message] struct {
	node        node.Node
	component   Component1[M]
	msgNotifyCh chan struct{}
}

func NewDataDrivenComponent1[M proto.Message](component Component1[M]) *DataDrivenComponent1[M] {
	return &DataDrivenComponent1[M]{
		component:   component,
		msgNotifyCh: make(chan struct{}),
	}
}

func (c DataDrivenComponent1[M]) Initialize(config *pb.ComponentConfig) bool {
	if len(config.GetName()) == 0 {
		log.Logger.Error("name empty")
		return false
	}

	if len(config.GetReaders()) < 1 {
		log.Logger.Error("invalid config file: too few readers")
		return false
	}

	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		log.Logger.Error("component init failed")
		return false
	}

	reader, err := node.NewChannelCreator[M](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[0].GetChannel(),
		QosProfile:       config.Readers[0].GetQosProfile(),
		PendingQueueSize: config.Readers[0].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	dv := data.NewDataVisitor[M](data.VistorConfig{
		ChannelID: reader.ChannelID(),
		QueueSize: reader.PendingQueueSize(),
	})

	dv.RegisterNotifyCallback(func() {
		select {
		case c.msgNotifyCh <- struct{}{}:
		default:
		}
	})

	go func() {
		for range c.msgNotifyCh {
			m := common.Zero[M]()
			if !dv.TryFetch(&m) {
				continue
			}
			c.Process(m)
		}
	}()
	return true
}

func (c *DataDrivenComponent1[M]) Process(m M) bool {
	return c.component.Proc(m)
}

type DataDrivenComponent2[M0, M1 proto.Message] struct {
	node        node.Node
	component   Component2[M0, M1]
	msgNotifyCh chan struct{}
}

func NewDataDrivenComponent2[M0, M1 proto.Message](component Component2[M0, M1]) *DataDrivenComponent2[M0, M1] {
	return &DataDrivenComponent2[M0, M1]{
		component:   component,
		msgNotifyCh: make(chan struct{}),
	}
}

func (c *DataDrivenComponent2[M0, M1]) Initialize(config *pb.ComponentConfig) bool {
	if len(config.GetName()) == 0 {
		log.Logger.Error("name empty")
		return false
	}

	if len(config.GetReaders()) < 2 {
		log.Logger.Error("invalid config file: too few readers")
		return false
	}
	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		log.Logger.Error("component init failed")
		return false
	}

	reader0, err := node.NewChannelCreator[M0](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[0].GetChannel(),
		QosProfile:       config.Readers[0].GetQosProfile(),
		PendingQueueSize: config.Readers[0].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader1, err := node.NewChannelCreator[M1](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[1].GetChannel(),
		QosProfile:       config.Readers[1].GetQosProfile(),
		PendingQueueSize: config.Readers[1].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	vistorConfigs := [2]data.VistorConfig{
		{
			ChannelID: reader0.ChannelID(),
			QueueSize: reader0.PendingQueueSize(),
		},
		{
			ChannelID: reader1.ChannelID(),
			QueueSize: reader1.PendingQueueSize(),
		},
	}

	dv := data.NewDataVisitor2[M0, M1](vistorConfigs)
	dv.RegisterNotifyCallback(func() {
		select {
		case c.msgNotifyCh <- struct{}{}:
		default:
		}
	})

	go func() {
		for range c.msgNotifyCh {
			m0 := common.Zero[M0]()
			m1 := common.Zero[M1]()
			if !dv.TryFetch(&m0, &m1) {
				continue
			}
			c.Process(m0, m1)
		}
	}()
	return true
}

func (c *DataDrivenComponent2[M0, M1]) Process(m0 M0, m1 M1) bool {
	return c.component.Proc(m0, m1)
}

type DataDrivenComponent3[M0, M1, M2 proto.Message] struct {
	node        node.Node
	component   Component3[M0, M1, M2]
	msgNotifyCh chan struct{}
}

func NewDataDrivenComponent3[M0, M1, M2 proto.Message](component Component3[M0, M1, M2]) *DataDrivenComponent3[M0, M1, M2] {
	return &DataDrivenComponent3[M0, M1, M2]{
		component:   component,
		msgNotifyCh: make(chan struct{}),
	}
}

func (c *DataDrivenComponent3[M0, M1, M2]) Initialize(config *pb.ComponentConfig) bool {
	if len(config.GetName()) == 0 {
		log.Logger.Error("name empty")
		return false
	}

	if len(config.GetReaders()) < 3 {
		log.Logger.Error("invalid config file: too few readers")
		return false
	}
	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		log.Logger.Error("component init failed")
		return false
	}

	reader0, err := node.NewChannelCreator[M0](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[0].GetChannel(),
		QosProfile:       config.Readers[0].GetQosProfile(),
		PendingQueueSize: config.Readers[0].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader1, err := node.NewChannelCreator[M1](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[1].GetChannel(),
		QosProfile:       config.Readers[1].GetQosProfile(),
		PendingQueueSize: config.Readers[1].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader2, err := node.NewChannelCreator[M2](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[2].GetChannel(),
		QosProfile:       config.Readers[2].GetQosProfile(),
		PendingQueueSize: config.Readers[2].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	vistorConfigs := [3]data.VistorConfig{
		{
			ChannelID: reader0.ChannelID(),
			QueueSize: reader0.PendingQueueSize(),
		},
		{
			ChannelID: reader1.ChannelID(),
			QueueSize: reader1.PendingQueueSize(),
		},
		{
			ChannelID: reader2.ChannelID(),
			QueueSize: reader2.PendingQueueSize(),
		},
	}

	dv := data.NewDataVisitor3[M0, M1, M2](vistorConfigs)
	dv.RegisterNotifyCallback(func() {
		select {
		case c.msgNotifyCh <- struct{}{}:
		default:
		}
	})

	go func() {
		for range c.msgNotifyCh {
			m0 := common.Zero[M0]()
			m1 := common.Zero[M1]()
			m2 := common.Zero[M2]()
			if !dv.TryFetch(&m0, &m1, &m2) {
				continue
			}
			c.Process(m0, m1, m2)
		}
	}()
	return true
}

func (c *DataDrivenComponent3[M0, M1, M2]) Process(m0 M0, m1 M1, m2 M2) bool {
	return c.component.Proc(m0, m1, m2)
}

type DataDrivenComponent4[M0, M1, M2, M3 proto.Message] struct {
	node        node.Node
	component   Component4[M0, M1, M2, M3]
	msgNotifyCh chan struct{}
}

func NewDataDrivenComponent4[M0, M1, M2, M3 proto.Message](component Component4[M0, M1, M2, M3]) *DataDrivenComponent4[M0, M1, M2, M3] {
	return &DataDrivenComponent4[M0, M1, M2, M3]{
		component:   component,
		msgNotifyCh: make(chan struct{}),
	}
}

func (c *DataDrivenComponent4[M0, M1, M2, M3]) Initialize(config *pb.ComponentConfig) bool {
	if len(config.GetName()) == 0 {
		log.Logger.Error("name empty")
		return false
	}

	if len(config.GetReaders()) < 4 {
		log.Logger.Error("invalid config file: too few readers")
		return false
	}
	c.node = node.NewChannelNode(config.GetName())

	if !c.component.Init(config.GetConfigFilePath(), c.node) {
		log.Logger.Error("component init failed")
		return false
	}

	reader0, err := node.NewChannelCreator[M0](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[0].GetChannel(),
		QosProfile:       config.Readers[0].GetQosProfile(),
		PendingQueueSize: config.Readers[0].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader1, err := node.NewChannelCreator[M1](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[1].GetChannel(),
		QosProfile:       config.Readers[1].GetQosProfile(),
		PendingQueueSize: config.Readers[1].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader2, err := node.NewChannelCreator[M2](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[2].GetChannel(),
		QosProfile:       config.Readers[2].GetQosProfile(),
		PendingQueueSize: config.Readers[2].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	reader3, err := node.NewChannelCreator[M3](c.node).CreateReaderByConfig(&node.ReaderConfig{
		ChannelName:      config.Readers[3].GetChannel(),
		QosProfile:       config.Readers[3].GetQosProfile(),
		PendingQueueSize: config.Readers[3].GetPendingQueueSize(),
	}, nil)
	if err != nil {
		log.Logger.Error("create reader failed", "err", err)
		return false
	}

	vistorConfigs := [4]data.VistorConfig{
		{
			ChannelID: reader0.ChannelID(),
			QueueSize: reader0.PendingQueueSize(),
		},
		{
			ChannelID: reader1.ChannelID(),
			QueueSize: reader1.PendingQueueSize(),
		},
		{
			ChannelID: reader2.ChannelID(),
			QueueSize: reader2.PendingQueueSize(),
		},
		{
			ChannelID: reader3.ChannelID(),
			QueueSize: reader3.PendingQueueSize(),
		},
	}

	dv := data.NewDataVisitor4[M0, M1, M2, M3](vistorConfigs)
	dv.RegisterNotifyCallback(func() {
		select {
		case c.msgNotifyCh <- struct{}{}:
		default:
		}
	})

	go func() {
		for range c.msgNotifyCh {
			m0 := common.Zero[M0]()
			m1 := common.Zero[M1]()
			m2 := common.Zero[M2]()
			m3 := common.Zero[M3]()
			if !dv.TryFetch(&m0, &m1, &m2, &m3) {
				continue
			}
			c.Process(m0, m1, m2, m3)
		}
	}()
	return true
}

func (c *DataDrivenComponent4[M0, M1, M2, M3]) Process(m0 M0, m1 M1, m2 M2, m3 M3) bool {
	return c.component.Proc(m0, m1, m2, m3)
}
