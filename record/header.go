package record

import (
	"fmt"
	"io"

	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type Header struct {
	SectionHeader *SectionHeader
	Header        *pb.Header
}

func (h Header) WriteTo(w io.Writer) (int64, error) {
	data, err := proto.Marshal(h.Header)
	if err != nil {
		return 0, err
	}

	sectionHeader := &SectionHeader{
		Size:        int64(len(data)),
		SectionType: pb.SectionType_SECTION_HEADER,
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

func NewHeaderFromReader(r io.Reader) (*Header, error) {
	sectionHeader, err := NewSectionHeaderFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
	}

	if sectionHeader.SectionType != pb.SectionType_SECTION_HEADER {
		return nil, fmt.Errorf("check section type failed, expect: %v, actual: %v", pb.SectionType_SECTION_HEADER,
			sectionHeader.SectionType)
	}

	reader := io.LimitReader(r, sectionHeader.Size)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var header pb.Header
	if err := proto.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	return &Header{
		SectionHeader: sectionHeader,
		Header:        &header,
	}, nil
}
