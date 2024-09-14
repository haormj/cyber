package record

import (
	"fmt"
	"io"

	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type Channel struct {
	SectionHeader *SectionHeader
	Channel       *pb.Channel
}

func (c Channel) WriteTo(w io.Writer) (int64, error) {
	data, err := proto.Marshal(c.Channel)
	if err != nil {
		return 0, err
	}

	sectionHeader := &SectionHeader{
		Size:        int64(len(data)),
		SectionType: pb.SectionType_SECTION_CHANNEL,
	}

	var written int64
	n, err := sectionHeader.WriteTo(w)
	if err != nil {
		return 0, err
	}
	written += n

	nn, err := w.Write(data)
	if err != nil {
		return 0, err
	}
	written += int64(nn)

	return written, nil
}

func NewChannelFromReader(r io.Reader) (*Channel, error) {
	sectionHeader, err := NewSectionHeaderFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
	}

	if sectionHeader.SectionType != pb.SectionType_SECTION_CHANNEL {
		return nil, fmt.Errorf("check section type failed, expect: %v, actual: %v", pb.SectionType_SECTION_CHANNEL,
			sectionHeader.SectionType)
	}

	reader := io.LimitReader(r, sectionHeader.Size)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var channel pb.Channel
	if err := proto.Unmarshal(data, &channel); err != nil {
		return nil, err
	}

	return &Channel{
		SectionHeader: sectionHeader,
		Channel:       &channel,
	}, nil
}
