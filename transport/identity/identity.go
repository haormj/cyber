package identity

import (
	"bytes"

	"github.com/google/uuid"
	"github.com/haormj/cyber/common"
	"github.com/spf13/cast"
)

const ID_SIZE uint8 = 8

type Identity struct {
	data      [ID_SIZE]byte
	hashValue uint64
}

func NewIdentity(needGenerate bool) *Identity {
	i := &Identity{}
	if needGenerate {
		u := [16]byte(uuid.New())
		copy(i.data[:], u[:8])
		i.update()
	}
	return i
}

func CloneIdentity(i *Identity) *Identity {
	ii := &Identity{
		hashValue: i.hashValue,
	}
	copy(ii.data[:], i.data[:])
	return ii
}

func (i *Identity) update() {
	i.hashValue = common.Hash(i.data[:])
}

func (i *Identity) ToString() string {
	return cast.ToString(i.hashValue)
}

func (i *Identity) Length() int {
	return int(ID_SIZE)
}

func (i *Identity) HashValue() uint64 {
	return i.hashValue
}

func (i *Identity) Data() []byte {
	return i.data[:]
}

func (i *Identity) SetData(data []byte) {
	if data == nil {
		return
	}

	copy(i.data[:], data)
	i.update()
}

func NotEqual(i1, i2 *Identity) bool {
	return !bytes.Equal(i1.Data(), i2.Data())
}
