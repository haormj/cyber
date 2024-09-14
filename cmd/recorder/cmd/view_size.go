package cmd

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/haormj/cyber/record"
	"github.com/haormj/util"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var viewSizeCmd = &cobra.Command{
	Use:   "size",
	Short: "recorder view size",
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

		root := util.NewNode()
		root.Data["name"] = t.Type
		for _, message := range messages {
			m, err := message.DynamicMessage()
			if err != nil {
				continue
			}

			root.Data["size"] = cast.ToInt(root.Data["size"]) + proto.Size(m)
			viewMessage(m.ProtoReflect(), 0, files, root, io.Discard)
		}

		total := cast.ToInt(root.Data["size"])
		root.DFSPreOrder(func(node *util.Node) {
			size := cast.ToInt(node.Data["size"])
			var ratio float64
			if total > 0 {
				ratio = float64(size) / float64(total) * 100
			}
			fmt.Printf("%s%s %d %.1f\n", strings.Repeat(" ", 2*node.Depth()), node.Data["name"], size, ratio)
		})
	},
}

func init() {
	viewSizeCmd.Flags().StringP("topic", "t", "", "topic name")
	if err := viewSizeCmd.MarkFlagRequired("topic"); err != nil {
		log.Fatalln(err)
	}
	viewCmd.AddCommand(viewSizeCmd)
}
