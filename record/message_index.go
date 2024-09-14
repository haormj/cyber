package record

import (
	"bytes"
	"compress/bzip2"
	"errors"
	"fmt"
	"io"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/btree"
	"github.com/haormj/cyber/pb"
	"github.com/pierrec/lz4/v4"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Topic struct {
	Name      string
	Type      string
	Files     *protoregistry.Files
	ProtoDesc []byte
}

func NewTopic(name string, message proto.Message) (*Topic, error) {
	pd, err := NewProtoDesc(message)
	if err != nil {
		return nil, fmt.Errorf("protodesc.NewProtoDesc %w", err)
	}

	dataType := string(message.ProtoReflect().Descriptor().FullName())

	topic := &Topic{
		Name:  name,
		Type:  dataType,
		Files: pd.RegistryFiles,
	}

	return topic, nil
}

var messageID uint64

func GenerateMessageID() uint64 {
	return atomic.AddUint64(&messageID, 1)
}

type Message struct {
	id    uint64
	Time  time.Time
	topic *Topic
	Data  []byte
}

func NewMessage(t time.Time, topic *Topic, m proto.Message) (*Message, error) {
	data, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	return NewMessageFromData(t, topic, data), nil
}

func NewMessageFromData(t time.Time, topic *Topic, data []byte) *Message {
	return &Message{
		id:    GenerateMessageID(),
		Time:  t,
		topic: topic,
		Data:  data,
	}
}

func (m *Message) DynamicMessage() (*dynamicpb.Message, error) {
	if m.topic == nil || m.topic.Files == nil {
		return nil, errors.New("can not find topic info")
	}

	d, err := m.topic.Files.FindDescriptorByName(protoreflect.FullName(m.topic.Type))
	if err != nil {
		return nil, fmt.Errorf("can not find %s", m.topic.Type)
	}

	md, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, errors.New("can't cast to protoreflect.MessageDescriptor")
	}

	dm := dynamicpb.NewMessage(dynamicpb.NewMessageType(md).Descriptor())
	if err := proto.Unmarshal(m.Data, dm); err != nil {
		return nil, err
	}
	return dm, nil
}

func (m *Message) To(v proto.Message) error {
	return proto.Unmarshal(m.Data, v)
}

func (m *Message) TopicName() string {
	return m.topic.Name
}

func (m *Message) DataLen() int {
	return len(m.Data)
}

func (m *Message) DeepCloneData() []byte {
	data := make([]byte, len(m.Data))
	copy(data, m.Data)
	return data
}

func (m *Message) CloneMessage() *Message {
	return &Message{
		id:    GenerateMessageID(),
		Time:  m.Time,
		topic: m.topic,
		Data:  m.DeepCloneData(),
	}
}

func (m *Message) ID() uint64 {
	return m.id
}

// MessageIndex 对topic和message建立相关index，方便操作
// 当前使用btree索引，对于两个相同的数据插入时会进行upsert，判断相关的标准就是传入的less
// 从而使用当前实现的btree索引，最好less中使用的字段能够唯一标识这条数据
type MessageIndex struct {
	lock             sync.RWMutex
	topicIndex       map[string]*Topic
	topicCount       map[string]int64
	messageTimeIndex *btree.BTreeG[*Message]
	messageNameIndex *btree.BTreeG[*Message]
}

func NewMessageIndex(topics []*Topic, messages []*Message, opts ...Option) (*MessageIndex, error) {
	var options Options
	for _, o := range opts {
		o(&options)
	}

	if options.TopicNameMapFunc != nil {
		for _, topic := range topics {
			topic.Name = options.TopicNameMapFunc(topic.Name)
		}
	}

	topicIndex := make(map[string]*Topic)
	topicCount := make(map[string]int64)

	messageTimeIndex := btree.NewG[*Message](10, func(a, b *Message) bool {
		if a.Time.Equal(b.Time) {
			if a.TopicName() == b.TopicName() {
				return a.ID() < b.ID()
			}

			return a.TopicName() < b.TopicName()
		}

		return a.Time.Before(b.Time)
	})

	messageNameIndex := btree.NewG[*Message](10, func(a, b *Message) bool {
		if a.TopicName() == b.TopicName() {
			if a.Time.Equal(b.Time) {
				return a.ID() < b.ID()
			}
			return a.Time.Before(b.Time)
		}

		return a.TopicName() < b.TopicName()
	})

	for _, topic := range topics {
		if topic == nil {
			continue
		}

		topicIndex[topic.Name] = topic
	}

	for _, message := range messages {
		if message == nil {
			continue
		}

		if _, ok := topicIndex[message.TopicName()]; !ok {
			return nil, errors.New("topic not exist")
		}

		topicCount[message.TopicName()] = topicCount[message.TopicName()] + 1
		messageTimeIndex.ReplaceOrInsert(message)
		messageNameIndex.ReplaceOrInsert(message)
	}

	messageIndex := &MessageIndex{
		topicIndex:       topicIndex,
		topicCount:       topicCount,
		messageTimeIndex: messageTimeIndex,
		messageNameIndex: messageNameIndex,
	}

	return messageIndex, nil
}

func (m *MessageIndex) TopicLen() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.topicIndex)
}

func (m *MessageIndex) AddTopic(topic *Topic) {
	if topic == nil {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.topicIndex[topic.Name] = topic
}

func (m *MessageIndex) GetTopicByName(name string) (*Topic, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	topic, ok := m.topicIndex[name]
	return topic, ok
}

func (m *MessageIndex) GetTopicFPS(name string) int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	count, ok := m.topicCount[name]
	if !ok {
		return 0
	}

	minMessage, _ := m.messageTimeIndex.Min()
	maxMessage, _ := m.messageTimeIndex.Max()

	if minMessage == nil || maxMessage == nil {
		return 0
	}

	seconds := maxMessage.Time.Sub(minMessage.Time).Seconds()
	if seconds == 0 {
		return 0
	}

	fps := math.Ceil(float64(count) / seconds)
	return int(fps)
}

type ListTopicsOptions struct {
	Names []string
}

type ListTopicsOption func(*ListTopicsOptions)

func ListTopicsWithName(names ...string) ListTopicsOption {
	return func(o *ListTopicsOptions) {
		o.Names = names
	}
}

func (m *MessageIndex) listTopics() []*Topic {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var topics []*Topic
	for _, topic := range m.topicIndex {
		topics = append(topics, topic)
	}

	return topics
}

// ListTopics 返回的topic list数组顺序是不稳定的
func (m *MessageIndex) ListTopics(opts ...ListTopicsOption) []*Topic {
	options := ListTopicsOptions{}
	for _, o := range opts {
		o(&options)
	}

	if len(options.Names) == 0 {
		return m.listTopics()
	}

	var topics []*Topic
	for _, name := range options.Names {
		t, ok := m.GetTopicByName(name)
		if ok {
			topics = append(topics, t)
		}
	}

	return topics
}

func (m *MessageIndex) MessageLen() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.messageTimeIndex.Len()
}

func (m *MessageIndex) MessageLenByName(name string) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.topicCount[name]
}

type ListMessagesOptions struct {
	Topics       []string
	Begin        *time.Time
	End          *time.Time
	MessageOrder messageOrder
}

type ListMessagesOption func(*ListMessagesOptions)

func ListMessagesWithTopics(topic ...string) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.Topics = topic
	}
}

func ListMessagesWithBegin(b time.Time) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.Begin = &b
	}
}

func ListMessagesWithEnd(e time.Time) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.End = &e
	}
}

type messageOrder int

const (
	MessageOrderAsc  messageOrder = 0
	MessageOrderDesc messageOrder = 1
)

func ListMessagesOrderBy(order messageOrder) ListMessagesOption {
	return func(o *ListMessagesOptions) {
		o.MessageOrder = order
	}
}

func (m *MessageIndex) ListMessages(opts ...ListMessagesOption) []*Message {
	options := ListMessagesOptions{}
	for _, o := range opts {
		o(&options)
	}

	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.messageTimeIndex.Len() == 0 {
		return nil
	}

	var begin, end time.Time
	if options.Begin != nil {
		begin = *options.Begin
	} else {
		minMessage, _ := m.messageTimeIndex.Min()
		begin = minMessage.Time
	}

	if options.End != nil {
		end = *options.End
	} else {
		maxMessage, _ := m.messageTimeIndex.Max()
		end = maxMessage.Time
	}
	// 因为btree这个包range时为 [a, b), 要实现 [a, b]，需要写为 [begin, end+1]
	end = end.Add(time.Nanosecond)

	var nameBegin, nameEnd string
	{
		minMessage, _ := m.messageNameIndex.Min()
		nameBegin = minMessage.TopicName()

		maxMessage, _ := m.messageNameIndex.Max()
		nameEnd = maxMessage.TopicName()
	}

	messageBegin := &Message{
		topic: &Topic{
			Name: nameBegin,
		},
		Time: begin,
	}

	messageEnd := &Message{
		topic: &Topic{
			Name: nameEnd,
		},
		Time: end,
	}

	var messages []*Message
	switch len(options.Topics) {
	case 0:
		switch options.MessageOrder {
		case MessageOrderAsc:
			m.messageTimeIndex.AscendRange(
				messageBegin,
				messageEnd,
				func(item *Message) bool {
					messages = append(messages, item)
					return true
				},
			)
		case MessageOrderDesc:
			m.messageTimeIndex.DescendRange(
				messageEnd,
				messageBegin,
				func(item *Message) bool {
					messages = append(messages, item)
					return true
				},
			)
		}
	default:
		for _, topic := range options.Topics {
			messageBegin.topic.Name = topic
			messageEnd.topic.Name = topic

			switch options.MessageOrder {
			case MessageOrderAsc:
				m.messageNameIndex.AscendRange(
					messageBegin,
					messageEnd,
					func(item *Message) bool {
						messages = append(messages, item)
						return true
					},
				)
			case MessageOrderDesc:
				m.messageNameIndex.DescendRange(
					messageEnd,
					messageBegin,
					func(item *Message) bool {
						messages = append(messages, item)
						return true
					},
				)
			}
		}
	}

	return messages
}

type CreateMessageIndexOptions struct {
	Topics []string
	Begin  *time.Time
	End    *time.Time
}

type CreateMessageIndexOption func(*CreateMessageIndexOptions)

func CreateMessageIndexWithTopic(topics ...string) CreateMessageIndexOption {
	return func(o *CreateMessageIndexOptions) {
		o.Topics = topics
	}
}

func CreateMessageIndexWithBegin(t time.Time) CreateMessageIndexOption {
	return func(o *CreateMessageIndexOptions) {
		o.Begin = &t
	}
}

func CreateMessageIndexWithEnd(t time.Time) CreateMessageIndexOption {
	return func(o *CreateMessageIndexOptions) {
		o.End = &t
	}
}

func (m *MessageIndex) CreateMessageIndex(opts ...CreateMessageIndexOption) (*MessageIndex, error) {
	options := CreateMessageIndexOptions{}
	for _, o := range opts {
		o(&options)
	}

	var listMessageOptions []ListMessagesOption
	if len(options.Topics) != 0 {
		listMessageOptions = append(listMessageOptions, ListMessagesWithTopics(options.Topics...))
	}
	if options.Begin != nil {
		listMessageOptions = append(listMessageOptions, ListMessagesWithBegin(*options.Begin))
	}
	if options.End != nil {
		listMessageOptions = append(listMessageOptions, ListMessagesWithEnd(*options.End))
	}

	messages := m.ListMessages(listMessageOptions...)

	topicMap := make(map[string]struct{})
	for _, message := range messages {
		topicMap[message.TopicName()] = struct{}{}
	}

	var topics []*Topic
	for k := range topicMap {
		topic, ok := m.GetTopicByName(k)
		if !ok {
			return nil, errors.New("topic not exist")
		}

		topics = append(topics, topic)
	}

	return NewMessageIndex(topics, messages)
}

func (m *MessageIndex) MinMessage() *Message {
	m.lock.RLock()
	defer m.lock.RUnlock()

	minMessage, _ := m.messageTimeIndex.Min()
	return minMessage
}

func (m *MessageIndex) MaxMessage() *Message {
	m.lock.RLock()
	defer m.lock.RUnlock()

	maxMessage, _ := m.messageTimeIndex.Max()
	return maxMessage
}

func (m *MessageIndex) AddMessage(message *Message) error {
	if message == nil {
		return nil
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.topicIndex[message.TopicName()]; !ok {
		return errors.New("topic not exist")
	}

	m.topicCount[message.TopicName()] = m.topicCount[message.TopicName()] + 1
	m.messageTimeIndex.ReplaceOrInsert(message)
	m.messageNameIndex.ReplaceOrInsert(message)

	return nil
}

// RepeatMessageIndex 寻找指定时刻的相关topic的消息(若未找到，则找该时刻前的第一个消息)
// 按照 step 向前重复，持续 duration
func (m *MessageIndex) RepeatMessageIndex(t time.Time, duration time.Duration,
	topicNames []string) (*MessageIndex, error) {
	topics := m.ListTopics(ListTopicsWithName(topicNames...))

	var messages []*Message
	for _, topic := range topics {
		var messsage *Message
		m.messageNameIndex.DescendLessOrEqual(
			&Message{
				topic: topic,
				Time:  t,
			},
			func(item *Message) bool {
				messsage = item
				return false
			},
		)

		if messsage == nil || messsage.TopicName() != topic.Name {
			continue
		}

		// 比目标时间差值小于 1s
		if t.Sub(messsage.Time) > time.Second {
			continue
		}

		fps := m.GetTopicFPS(topic.Name)
		if fps == 0 {
			continue
		}

		step := time.Second / time.Duration(fps)
		for i := 0; ; i++ {
			cloneMessage := messsage.CloneMessage()
			cloneMessage.Time = cloneMessage.Time.Add(time.Duration(i) * step)
			messages = append(messages, cloneMessage)

			if cloneMessage.Time.Sub(t) > duration {
				break
			}
		}
	}

	return NewMessageIndex(topics, messages)
}
func NewMessageIndexFromPath(p string, opts ...Option) (*MessageIndex, error) {
	r, err := DecodeFromPath(p)
	if err != nil {
		return nil, fmt.Errorf("record.DecodeFromPath: %w", err)
	}
	return NewMessageIndexFromRecord(r, opts...)
}

func MergeMessageIndex(mis ...*MessageIndex) (*MessageIndex, error) {
	if len(mis) == 0 {
		return nil, errors.New("no MessageIndex need to merge")
	}

	var topics []*Topic
	var messages []*Message
	for _, mi := range mis {
		if mi == nil {
			continue
		}

		topics = append(topics, mi.ListTopics()...)
		messages = append(messages, mi.ListMessages()...)
	}

	return NewMessageIndex(topics, messages)
}

func NewMessageIndexFromRecord(r *Record, opts ...Option) (*MessageIndex, error) {
	topicMap := make(map[string]*Topic)

	newTopicFromChannel := func(channel *pb.Channel) *Topic {
		files := new(protoregistry.Files)
		pd, err := NewProtoDescFromBytes(channel.ProtoDesc)
		if err == nil {
			files = pd.RegistryFiles
		}

		topic := &Topic{
			Name:      channel.GetName(),
			Type:      channel.GetMessageType(),
			Files:     files,
			ProtoDesc: channel.ProtoDesc,
		}

		return topic
	}

	for _, channel := range r.Channels {
		topicMap[channel.Channel.GetName()] = newTopicFromChannel(channel.Channel)
	}

	var messages []*Message
	for _, chunkBody := range r.ChunkBodies {
		for _, msg := range chunkBody.ChunkBody.Messages {
			if msg == nil {
				continue
			}

			topic, ok := topicMap[msg.GetChannelName()]
			if !ok {
				return nil, fmt.Errorf("can not find topic %s", msg.GetChannelName())
			}

			var reader io.Reader
			reader = bytes.NewReader(msg.Content)
			switch r.Header.Header.GetCompress() {
			case pb.CompressType_COMPRESS_BZ2:
				reader = bzip2.NewReader(reader)
			case pb.CompressType_COMPRESS_LZ4:
				reader = lz4.NewReader(reader)
			case pb.CompressType_COMPRESS_NONE:
			default:
				return nil, fmt.Errorf("can not find compress_type %s", r.Header.Header.GetCompress())
			}

			data, err := io.ReadAll(reader)
			if err != nil {
				return nil, fmt.Errorf("io.ReadAll: %w", err)
			}

			messages = append(messages, &Message{
				id:    GenerateMessageID(),
				Time:  time.Unix(int64(msg.GetTime())/1e9, int64(msg.GetTime())%1e9),
				topic: topic,
				Data:  data,
			})
		}
	}

	var topics []*Topic
	for _, v := range topicMap {
		topics = append(topics, v)
	}

	return NewMessageIndex(topics, messages, opts...)
}
