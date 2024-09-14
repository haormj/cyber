package record

import (
	"fmt"
	"io"

	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type Index struct {
	SectionHeader *SectionHeader
	IndexSection  *pb.Index
}

func (i Index) WriteTo(w io.Writer) (int64, error) {
	data, err := proto.Marshal(i.IndexSection)
	if err != nil {
		return 0, err
	}

	sectionHeader := &SectionHeader{
		Size:        int64(len(data)),
		SectionType: pb.SectionType_SECTION_INDEX,
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

func NewIndexFromReader(r io.ReadSeeker, header *Header) (*Index, error) {
	if _, err := r.Seek(int64(header.Header.GetIndexPosition()), io.SeekStart); err != nil {
		return nil, err
	}

	sectionHeader, err := NewSectionHeaderFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
	}

	if sectionHeader.SectionType != pb.SectionType_SECTION_INDEX {
		return nil, fmt.Errorf("check section type failed, expect: %v, actual: %v", pb.SectionType_SECTION_INDEX,
			sectionHeader.SectionType)
	}

	reader := io.LimitReader(r, sectionHeader.Size)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}
	var indexSection pb.Index
	if err := proto.Unmarshal(data, &indexSection); err != nil {
		return nil, fmt.Errorf("proto.Unmarshal: %w", err)
	}

	return &Index{
		SectionHeader: sectionHeader,
		IndexSection:  &indexSection,
	}, nil
}
