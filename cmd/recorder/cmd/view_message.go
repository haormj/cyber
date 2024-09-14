package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/haormj/cyber/record"
	"github.com/haormj/util"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

var viewMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "recorder view message",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must have one record path")
		}

		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			log.Fatalln(err)
		}

		if len(topic) == 0 {
			log.Fatalln("must have topic and topic not empty")
		}

		begin, err := cmd.Flags().GetString("begin")
		if err != nil {
			log.Fatalln(err)
		}

		end, err := cmd.Flags().GetString("end")
		if err != nil {
			log.Fatalln(err)
		}

		var b int64 = 0
		if begin != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05.999999999", begin, time.Local)
			if err != nil {
				log.Fatalln(err)
			}
			b = t.UnixNano()
		}

		var e int64 = math.MaxInt64
		if end != "" {
			t, err := time.ParseInLocation("2006-01-02 15:04:05.999999999", end, time.Local)
			if err != nil {
				log.Fatalln(err)
			}
			e = t.UnixNano()
		}

		if e < b {
			log.Fatalln("begin must < end")
		}

		messageIndex, err := record.NewMessageIndexFromPath(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		files := protoregistry.GlobalFiles
		for _, topic := range messageIndex.ListTopics() {
			topic.Files.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
				if _, err := files.FindFileByPath(fd.Path()); err == nil {
					return true
				}

				if err == protoregistry.NotFound {
					_ = files.RegisterFile(fd)
				}

				return true
			})
		}

		mi, err := messageIndex.CreateMessageIndex(
			record.CreateMessageIndexWithTopic(topic),
			record.CreateMessageIndexWithBegin(time.Unix(b/1e9, b%1e9)),
			record.CreateMessageIndexWithEnd(time.Unix(e/1e9, e%1e9)),
		)
		if err != nil {
			log.Fatalln(err)
		}

		ts := mi.ListTopics()
		if len(ts) != 1 {
			log.Fatalln("len(topics) != 1")
		}
		t := ts[0]

		messages := mi.ListMessages()
		fmt.Println("topic:", t.Name)
		fmt.Println("type:", t.Type)
		fmt.Println("count: ", len(messages))

		fmt.Println("messages: ")
		root := util.NewNode()
		root.Data["name"] = t.Type
		for _, message := range messages {
			fmt.Println(message.Time.Format("2006-01-02 15:04:05.999999999"))
			m, err := message.DynamicMessage()
			if err != nil {
				fmt.Println("  decode message failed", err)
				continue
			}
			root.Data["size"] = cast.ToInt(root.Data["size"]) + proto.Size(m)
			viewMessage(m.ProtoReflect(), 0, files, root, os.Stdout)
		}
	},
}

func init() {
	viewMessageCmd.Flags().StringP("topic", "t", "", "topic name")
	if err := viewMessageCmd.MarkFlagRequired("topic"); err != nil {
		log.Fatalln(err)
	}

	viewMessageCmd.Flags().StringP("begin", "b", "", "begin time(local timezone), date format 2006-01-02 15:04:05.999999999")
	viewMessageCmd.Flags().StringP("end", "e", "", "end time(local timezone), date format 2006-01-02 15:04:05.999999999")

	viewCmd.AddCommand(viewMessageCmd)
}

func castAnyToMessage(files *protoregistry.Files, message protoreflect.Message) (protoreflect.Message, error) {
	if message.Descriptor().FullName() != "google.protobuf.Any" {
		return message, nil
	}

	typeURL := message.Get(message.Descriptor().Fields().ByName("type_url")).String()
	value := message.Get(message.Descriptor().Fields().ByName("value")).Bytes()

	d, err := files.FindDescriptorByName(protoreflect.FullName(strings.Split(typeURL, "/")[1]))
	if err != nil {
		return nil, err
	}

	md, ok := d.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, errors.New("can't cast to protoreflect.MessageDescriptor")
	}

	m := dynamicpb.NewMessage(dynamicpb.NewMessageType(md).Descriptor())
	if err := proto.Unmarshal(value, m); err != nil {
		return nil, err
	}

	return m, nil
}

func viewMessage(message protoreflect.Message, depth int, files *protoregistry.Files,
	node *util.Node, writer io.Writer) {
	prefix := strings.Repeat(" ", 2*(depth+1))

	if message.Descriptor().FullName() == "google.protobuf.Any" {
		m, err := castAnyToMessage(files, message)
		if err == nil {
			typeURL := message.Get(message.Descriptor().Fields().ByName("type_url")).String()
			fmt.Fprintln(writer, prefix, "type_url", "string", typeURL)
			viewMessage(m, depth, files, node, writer)
			return
		}
	}

	message.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		n, ok := node.Search(
			"name",
			func(v any, value any) bool {
				return v.(string) == value.(string)
			},
			string(fd.Name()),
		)
		if !ok {
			n = util.NewNode()
			n.Data["name"] = string(fd.Name())
			node.AddChild(n)
		}
		n.Data["size"] = cast.ToInt(n.Data["size"]) + sizeField(fd, v)

		if fd.Cardinality() == protoreflect.Repeated {
			switch {
			case fd.IsList():
				fmt.Fprintf(writer, "%s %s []%s\n", prefix, fd.Name(), fd.Kind())

				l := v.List()
				for i := 0; i < l.Len(); i++ {
					fmt.Fprintln(writer, prefix+"  ", i, l.Get(i))
				}
			case fd.IsMap():
				fmt.Fprintf(writer, "%s %s %s map[%s]%s\n", prefix, fd.Name(), fd.Kind(),
					fd.MapKey().Kind(), fd.MapValue().Kind())

				mapValue := v.Map()
				mapValue.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
					fmt.Fprintf(writer, "%s  %s: %s\n", prefix, k, v)
					return true
				})
			default:
				fmt.Fprintln(writer, prefix, fd.Name(), fd.Kind(), "repeated", "no support decode")
			}

			return true
		}

		if fd.Kind() != protoreflect.MessageKind {
			if fd.Kind() == protoreflect.BytesKind {
				fmt.Fprintln(writer, prefix, fd.Name(), fd.Kind(), len(v.Bytes()))
			} else {
				fmt.Fprintln(writer, prefix, fd.Name(), fd.Kind(), v)
			}
			return true
		}

		fmt.Fprintln(writer, prefix, fd.Name(), fd.Kind(), fd.FullName())
		viewMessage(v.Message(), depth+1, files, n, writer)
		return true
	})
}

// IsMessageSet returns whether the message uses the MessageSet wire format.
func IsMessageSet(md protoreflect.MessageDescriptor) bool {
	xmd, ok := md.(interface{ IsMessageSet() bool })
	return ok && xmd.IsMessageSet()
}

func sizeMessageSlow(m protoreflect.Message) (size int) {
	if IsMessageSet(m.Descriptor()) {
		panic("not support message set")
	}
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += sizeField(fd, v)
		return true
	})
	size += len(m.GetUnknown())
	return size
}

func sizeField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (size int) {
	num := fd.Number()
	switch {
	case fd.IsList():
		return sizeList(num, fd, value.List())
	case fd.IsMap():
		return sizeMap(num, fd, value.Map())
	default:
		return protowire.SizeTag(num) + sizeSingular(num, fd.Kind(), value)
	}
}

func sizeList(num protowire.Number, fd protoreflect.FieldDescriptor, list protoreflect.List) (size int) {
	if fd.IsPacked() && list.Len() > 0 {
		content := 0
		for i, llen := 0, list.Len(); i < llen; i++ {
			content += sizeSingular(num, fd.Kind(), list.Get(i))
		}
		return protowire.SizeTag(num) + protowire.SizeBytes(content)
	}

	for i, llen := 0, list.Len(); i < llen; i++ {
		size += protowire.SizeTag(num) + sizeSingular(num, fd.Kind(), list.Get(i))
	}
	return size
}

func sizeMap(num protowire.Number, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) (size int) {
	mapv.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		size += protowire.SizeTag(num)
		size += protowire.SizeBytes(sizeField(fd.MapKey(), key.Value()) + sizeField(fd.MapValue(), value))
		return true
	})
	return size
}

func sizeSingular(num protowire.Number, kind protoreflect.Kind, v protoreflect.Value) int {
	switch kind {
	case protoreflect.BoolKind:
		return protowire.SizeVarint(protowire.EncodeBool(v.Bool()))
	case protoreflect.EnumKind:
		return protowire.SizeVarint(uint64(v.Enum()))
	case protoreflect.Int32Kind:
		return protowire.SizeVarint(uint64(int32(v.Int())))
	case protoreflect.Sint32Kind:
		return protowire.SizeVarint(protowire.EncodeZigZag(int64(int32(v.Int()))))
	case protoreflect.Uint32Kind:
		return protowire.SizeVarint(uint64(uint32(v.Uint())))
	case protoreflect.Int64Kind:
		return protowire.SizeVarint(uint64(v.Int()))
	case protoreflect.Sint64Kind:
		return protowire.SizeVarint(protowire.EncodeZigZag(v.Int()))
	case protoreflect.Uint64Kind:
		return protowire.SizeVarint(v.Uint())
	case protoreflect.Sfixed32Kind:
		return protowire.SizeFixed32()
	case protoreflect.Fixed32Kind:
		return protowire.SizeFixed32()
	case protoreflect.FloatKind:
		return protowire.SizeFixed32()
	case protoreflect.Sfixed64Kind:
		return protowire.SizeFixed64()
	case protoreflect.Fixed64Kind:
		return protowire.SizeFixed64()
	case protoreflect.DoubleKind:
		return protowire.SizeFixed64()
	case protoreflect.StringKind:
		return protowire.SizeBytes(len(v.String()))
	case protoreflect.BytesKind:
		return protowire.SizeBytes(len(v.Bytes()))
	case protoreflect.MessageKind:
		return protowire.SizeBytes(sizeMessageSlow(v.Message()))
	case protoreflect.GroupKind:
		return protowire.SizeGroup(num, sizeMessageSlow(v.Message()))
	default:
		return 0
	}
}
