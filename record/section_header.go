package record

import (
	"encoding/binary"
	"io"

	"github.com/haormj/cyber/pb"
)

type SectionHeader struct {
	SectionType pb.SectionType
	padding     int32
	Size        int64
}

func (h SectionHeader) WriteTo(w io.Writer) (int64, error) {
	if err := binary.Write(w, binary.LittleEndian, h.SectionType); err != nil {
		return 0, err
	}

	var padding int32
	if err := binary.Write(w, binary.LittleEndian, padding); err != nil {
		return 0, err
	}

	if err := binary.Write(w, binary.LittleEndian, h.Size); err != nil {
		return 0, err
	}

	return 8 + 4 + 4, nil
}

func NewSectionHeaderFromReader(f io.Reader) (*SectionHeader, error) {
	var sectionType pb.SectionType
	if err := binary.Read(f, binary.LittleEndian, &sectionType); err != nil {
		return nil, err
	}

	var padding int32
	if err := binary.Read(f, binary.LittleEndian, &padding); err != nil {
		return nil, err
	}

	var size int64
	if err := binary.Read(f, binary.LittleEndian, &size); err != nil {
		return nil, err
	}

	return &SectionHeader{
		SectionType: sectionType,
		padding:     padding,
		Size:        size,
	}, nil
}
