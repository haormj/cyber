package record

import (
	"fmt"
	"io"

	"github.com/haormj/cyber/pb"
	"google.golang.org/protobuf/proto"
)

type ChunkHeader struct {
	SectionHeader *SectionHeader
	ChunkHeader   *pb.ChunkHeader
}

func (c ChunkHeader) WriteTo(w io.Writer) (int64, error) {
	data, err := proto.Marshal(c.ChunkHeader)
	if err != nil {
		return 0, err
	}

	sectionHeader := &SectionHeader{
		Size:        int64(len(data)),
		SectionType: pb.SectionType_SECTION_CHUNK_HEADER,
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

func NewChunkHeaderFromReader(r io.Reader) (*ChunkHeader, error) {
	sectionHeader, err := NewSectionHeaderFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
	}

	if sectionHeader.SectionType != pb.SectionType_SECTION_CHUNK_HEADER {
		return nil, fmt.Errorf("check section type failed, expect: %v, actual: %v", pb.SectionType_SECTION_CHUNK_HEADER,
			sectionHeader.SectionType)
	}

	reader := io.LimitReader(r, sectionHeader.Size)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var chunkHeader pb.ChunkHeader
	if err := proto.Unmarshal(data, &chunkHeader); err != nil {
		return nil, err
	}

	return &ChunkHeader{
		SectionHeader: sectionHeader,
		ChunkHeader:   &chunkHeader,
	}, nil
}
