package common

import (
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

func GetProtoFromFile(p string, m proto.Message) error {
	b, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	return prototext.Unmarshal(b, m)
}

func GetFileName(p string) string {
	name := filepath.Base(p)
	ext := filepath.Ext(name)
	return strings.TrimRight(name, ext)
}
