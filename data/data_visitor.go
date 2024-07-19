package data

type VistorConfig struct {
	ChannelID uint64
	QueueSize uint32
}

type DataVisitor[T any] struct {
	notifier     *Notifier
	dataNotifier *DataNotifier
	buffer       *ChannelBuffer[T]
	nextMsgIndex uint64
}

func NewDataVisitor[T any](config VistorConfig) *DataVisitor[T] {
	buffer := NewChannelBuffer[T](config.ChannelID, NewCacheBuffer[T](int(config.QueueSize)))
	AddBuffer(buffer)
	v := &DataVisitor[T]{
		notifier:     &Notifier{},
		dataNotifier: DataNotifierInstance,
		buffer:       buffer,
	}
	v.dataNotifier.AddNotifier(config.ChannelID, v.notifier)

	return v
}

func (v *DataVisitor[T]) RegisterNotifyCallback(callback func()) {
	v.notifier.Callback = callback
}

func (v *DataVisitor[T]) TryFetch(m *T) bool {
	if v.buffer.Fetch(&v.nextMsgIndex, m) {
		v.nextMsgIndex++
		return true
	}
	return false
}

type DataVisitor2[T0, T1 any] struct {
	notifier     *Notifier
	dataNotifier *DataNotifier
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	dataFusion   DataFusion2[T0, T1]
	nextMsgIndex uint64
}

func NewDataVisitor2[T0, T1 any](config [2]VistorConfig) *DataVisitor2[T0, T1] {
	buffer0 := NewChannelBuffer[T0](config[0].ChannelID, NewCacheBuffer[T0](int(config[0].QueueSize)))
	buffer1 := NewChannelBuffer[T1](config[1].ChannelID, NewCacheBuffer[T1](int(config[1].QueueSize)))
	AddBuffer(buffer0)
	AddBuffer(buffer1)
	v := &DataVisitor2[T0, T1]{
		notifier:     &Notifier{},
		dataNotifier: DataNotifierInstance,
		buffer0:      buffer0,
		buffer1:      buffer1,
	}
	v.dataNotifier.AddNotifier(config[0].ChannelID, v.notifier)
	v.dataFusion = NewAllLatest2[T0, T1](buffer0, buffer1)
	return v
}

func (v *DataVisitor2[T0, T1]) TryFetch(t0 *T0, t1 *T1) bool {
	if v.dataFusion.Fusion(&v.nextMsgIndex, t0, t1) {
		v.nextMsgIndex++
		return true
	}
	return false
}

func (v *DataVisitor2[T0, T1]) RegisterNotifyCallback(callback func()) {
	v.notifier.Callback = callback
}

type DataVisitor3[T0, T1, T2 any] struct {
	notifier     *Notifier
	dataNotifier *DataNotifier
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	buffer2      *ChannelBuffer[T2]
	dataFusion   DataFusion3[T0, T1, T2]
	nextMsgIndex uint64
}

func NewDataVisitor3[T0, T1, T2 any](config [3]VistorConfig) *DataVisitor3[T0, T1, T2] {
	buffer0 := NewChannelBuffer[T0](config[0].ChannelID, NewCacheBuffer[T0](int(config[0].QueueSize)))
	buffer1 := NewChannelBuffer[T1](config[1].ChannelID, NewCacheBuffer[T1](int(config[1].QueueSize)))
	buffer2 := NewChannelBuffer[T2](config[2].ChannelID, NewCacheBuffer[T2](int(config[2].QueueSize)))
	AddBuffer(buffer0)
	AddBuffer(buffer1)
	AddBuffer(buffer2)
	v := &DataVisitor3[T0, T1, T2]{
		notifier:     &Notifier{},
		dataNotifier: DataNotifierInstance,
		buffer0:      buffer0,
		buffer1:      buffer1,
		buffer2:      buffer2,
	}
	v.dataNotifier.AddNotifier(config[0].ChannelID, v.notifier)
	v.dataFusion = NewAllLatest3[T0, T1, T2](buffer0, buffer1, buffer2)
	return v
}

func (v *DataVisitor3[T0, T1, T2]) TryFetch(t0 *T0, t1 *T1, t2 *T2) bool {
	if v.dataFusion.Fusion(&v.nextMsgIndex, t0, t1, t2) {
		v.nextMsgIndex++
		return true
	}
	return false
}

func (v *DataVisitor3[T0, T1, T2]) RegisterNotifyCallback(callback func()) {
	v.notifier.Callback = callback
}

type DataVisitor4[T0, T1, T2, T3 any] struct {
	notifier     *Notifier
	dataNotifier *DataNotifier
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	buffer2      *ChannelBuffer[T2]
	buffer3      *ChannelBuffer[T3]
	dataFusion   DataFusion4[T0, T1, T2, T3]
	nextMsgIndex uint64
}

func NewDataVisitor4[T0, T1, T2, T3 any](config [4]VistorConfig) *DataVisitor4[T0, T1, T2, T3] {
	buffer0 := NewChannelBuffer[T0](config[0].ChannelID, NewCacheBuffer[T0](int(config[0].QueueSize)))
	buffer1 := NewChannelBuffer[T1](config[1].ChannelID, NewCacheBuffer[T1](int(config[1].QueueSize)))
	buffer2 := NewChannelBuffer[T2](config[2].ChannelID, NewCacheBuffer[T2](int(config[2].QueueSize)))
	buffer3 := NewChannelBuffer[T3](config[3].ChannelID, NewCacheBuffer[T3](int(config[3].QueueSize)))
	AddBuffer(buffer0)
	AddBuffer(buffer1)
	AddBuffer(buffer2)
	AddBuffer(buffer3)
	v := &DataVisitor4[T0, T1, T2, T3]{
		notifier:     &Notifier{},
		dataNotifier: DataNotifierInstance,
		buffer0:      buffer0,
		buffer1:      buffer1,
		buffer2:      buffer2,
		buffer3:      buffer3,
	}
	v.dataNotifier.AddNotifier(config[0].ChannelID, v.notifier)
	v.dataFusion = NewAllLatest4[T0, T1, T2, T3](buffer0, buffer1, buffer2, buffer3)
	return v
}

func (v *DataVisitor4[T0, T1, T2, T3]) TryFetch(t0 *T0, t1 *T1, t2 *T2, t3 *T3) bool {
	if v.dataFusion.Fusion(&v.nextMsgIndex, t0, t1, t2, t3) {
		v.nextMsgIndex++
		return true
	}
	return false
}

func (v *DataVisitor4[T0, T1, T2, T3]) RegisterNotifyCallback(callback func()) {
	v.notifier.Callback = callback
}
