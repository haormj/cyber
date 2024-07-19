package common

import (
	"reflect"
)

func Hash(key []byte) uint64 {
	return MurmurHash64A(key, 0xc70f6907)
}

// MurmurHash64A 是从 github.com/rryqszq4/go-murmurhash 复制而来
// 目前测试下来与 cyber std::hash 中的结果一致
// 二者使用的都是 murmur2
func MurmurHash64A(key []byte, seed uint64) (hash uint64) {
	const m uint64 = 0xc6a4a7935bd1e995
	const r = 47

	var l int = len(key)
	var h uint64 = seed ^ uint64(l)*m

	var data []byte = key
	var l8 int = l / 8

	var k uint64

	for i := 0; i < l8; i++ {
		i8 := i * 8
		k = uint64(data[i8+0]) + uint64(data[i8+1])<<8 +
			uint64(data[i8+2])<<16 + uint64(data[i8+3])<<24 +
			uint64(data[i8+4])<<32 + uint64(data[i8+5])<<40 +
			uint64(data[i8+6])<<48 + uint64(data[i8+7])<<56

		k *= m
		k ^= k >> r
		k *= m

		h ^= k
		h *= m
	}

	data = data[l8*8:]

	switch l & 7 {
	case 7:
		h ^= uint64(data[6]) << 48
		fallthrough
	case 6:
		h ^= uint64(data[5]) << 40
		fallthrough
	case 5:
		h ^= uint64(data[4]) << 32
		fallthrough
	case 4:
		h ^= uint64(data[3]) << 24
		fallthrough
	case 3:
		h ^= uint64(data[2]) << 16
		fallthrough
	case 2:
		h ^= uint64(data[1]) << 8
		fallthrough
	case 1:
		h ^= uint64(data[0])
		h *= m
	}

	h ^= h >> r
	h *= m
	h ^= h >> r

	return h
}

func Zero[T any]() T {
	return ZeroByType(reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
}

func ZeroByType(typ reflect.Type) reflect.Value {
	if typ.Kind() == reflect.Ptr {
		data := ZeroByType(typ.Elem())
		v := reflect.New(typ.Elem())
		if v.CanSet() && data.CanAddr() {
			v.Set(data.Addr())
		}
		return v
	}

	return reflect.Zero(typ)
}
