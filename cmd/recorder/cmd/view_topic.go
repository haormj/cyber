package cmd

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/haormj/cyber/record"
	"github.com/spf13/cobra"
)

var viewTopicCmd = &cobra.Command{
	Use:   "topic",
	Short: "recorder view topic",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must have one record path")
		}

		fileInfo, err := os.Stat(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		total := fileInfo.Size()
		messageIndex, err := record.NewMessageIndexFromPath(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		var duration time.Duration
		if messageIndex.MessageLen() != 0 {
			duration = messageIndex.MaxMessage().Time.Sub(messageIndex.MinMessage().Time)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprint(writer, "name\ttype\tcount\tfps\tsize(byte)\tsize_ratio\tproto_desc\n")
		topics := messageIndex.ListTopics()
		var topicsNames []string
		for _, topic := range topics {
			topicsNames = append(topicsNames, topic.Name)
		}
		sort.Strings(topicsNames)

		for _, topicName := range topicsNames {
			topic, ok := messageIndex.GetTopicByName(topicName)
			if !ok {
				continue
			}

			messages := messageIndex.ListMessages(record.ListMessagesWithTopics(topic.Name))
			if len(messages) == 0 {
				fmt.Fprintf(writer, "%s\t%s\t%d\t0\t0\t0.0\t%d\n", topic.Name, topic.Type, len(messages),
					len(topic.ProtoDesc))
				continue
			}

			var fps float64
			if duration > 0 {
				fps = math.Ceil((float64(len(messages))) / duration.Seconds())
			}

			var size int64
			for _, message := range messages {
				size += int64(message.DataLen())
			}
			var sizeRatio float64
			if total > 0 {
				sizeRatio = float64(size) / float64(total) * 100
			}

			fmt.Fprintf(writer, "%s\t%s\t%d\t%.0f\t%d\t%.1f\t%d\n", topic.Name, topic.Type, len(messages),
				fps, size, sizeRatio, len(topic.ProtoDesc))
		}

		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "topic total\t%d\n", messageIndex.TopicLen())
		fmt.Fprintf(writer, "message total\t%d\n", messageIndex.MessageLen())
		if messageIndex.MessageLen() != 0 {
			fmt.Fprintf(writer, "start\t%v\n", messageIndex.MinMessage().Time)
			fmt.Fprintf(writer, "end\t%v\n", messageIndex.MaxMessage().Time)
		}
		fmt.Fprintf(writer, "duration\t%s\n", duration)
		fmt.Fprintf(writer, "size(byte)\t%d\n", total)

		if err := writer.Flush(); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	viewCmd.AddCommand(viewTopicCmd)
}
