package record

import (
	"fmt"
	"io"
	"os"

	"github.com/haormj/cyber/pb"
)

type Record struct {
	Header       *Header
	Index        *Index
	Channels     []*Channel
	ChunkHeaders []*ChunkHeader
	ChunkBodies  []*ChunkBody
}

func DecodeFromPath(path string) (*Record, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer f.Close()

	b, err := DecodeFromFile(f)
	if err != nil {
		return nil, fmt.Errorf("DecodeFromFile: %w", err)
	}

	return b, nil
}

func DecodeFromFile(f *os.File) (*Record, error) {
	b, err := DecodeFromReader(f)
	if err != nil {
		return nil, fmt.Errorf("DecodeFromReader: %w", err)
	}

	return b, nil
}

func DecodeFromReader(f io.ReadSeeker) (*Record, error) {
	header, err := NewHeaderFromReader(f)
	if err != nil {
		return nil, fmt.Errorf("NewHeaderFromReader: %w", err)
	}

	index, err := NewIndexFromReader(f, header)
	if err != nil {
		return nil, fmt.Errorf("NewIndexFromReader: %w", err)
	}

	var channels []*Channel
	var chunkHeaders []*ChunkHeader
	var chunkBodies []*ChunkBody
IndexLoop:
	for _, index := range index.IndexSection.Indexes {
		if _, err := f.Seek(int64(index.GetPosition()), io.SeekStart); err != nil {
			return nil, err
		}

		sectionHeader, err := NewSectionHeaderFromReader(f)
		if err != nil {
			return nil, fmt.Errorf("NewSectionHeaderFromReader: %w", err)
		}

		if _, err := f.Seek(int64(index.GetPosition()), io.SeekStart); err != nil {
			return nil, err
		}

		switch sectionHeader.SectionType {
		case pb.SectionType_SECTION_CHANNEL:
			channel, err := NewChannelFromReader(f)
			if err != nil {
				return nil, fmt.Errorf("NewChannelFromReader: %w", err)
			}
			channels = append(channels, channel)
		case pb.SectionType_SECTION_CHUNK_HEADER:
			chunkHeader, err := NewChunkHeaderFromReader(f)
			if err != nil {
				return nil, fmt.Errorf("NewChunkHeaderFromReader: %w", err)
			}
			chunkHeaders = append(chunkHeaders, chunkHeader)
		case pb.SectionType_SECTION_CHUNK_BODY:
			chunkBody, err := NewChunkBodyFromReader(f)
			if err != nil {
				return nil, fmt.Errorf("NewChunkBodyFromReader: %w", err)
			}
			chunkBodies = append(chunkBodies, chunkBody)
		case pb.SectionType_SECTION_INDEX:
			break IndexLoop
		default:
			return nil, fmt.Errorf("not support %v", sectionHeader.SectionType)
		}
	}

	return &Record{
		Header:       header,
		Index:        index,
		Channels:     channels,
		ChunkHeaders: chunkHeaders,
		ChunkBodies:  chunkBodies,
	}, nil
}
