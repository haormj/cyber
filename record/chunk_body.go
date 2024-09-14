package record

import (
	"fmt"
	"io"

	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type ChunkBody struct {
	SectionHeader *SectionHeader
	ChunkBody     *pb.ChunkBody
}

func (c ChunkBody) WriteTo(w io.Writer) (int64, error) {
	data, err := proto.Marshal(c.ChunkBody)
	if err != nil {
		return 0, err
	}

	sectionHeader := &SectionHeader{
		Size:        int64(len(data)),
		SectionType: pb.SectionType_SECTION_CHUNK_BODY,
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

func NewChunkBodyFromReader(r io.Reader) (*ChunkBody, error) {
	sectionHeader, err := NewSectionHeaderFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
	}

	if sectionHeader.SectionType != pb.SectionType_SECTION_CHUNK_BODY {
		return nil, fmt.Errorf("check section type failed, expect: %v, actual: %v", pb.SectionType_SECTION_CHUNK_BODY,
			sectionHeader.SectionType)
	}

	reader := io.LimitReader(r, sectionHeader.Size)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var body pb.ChunkBody
	if err := proto.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return &ChunkBody{
		SectionHeader: sectionHeader,
		ChunkBody:     &body,
	}, nil
}
