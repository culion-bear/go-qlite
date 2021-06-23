package node

import (
	"bytes"
	"hash/crc32"
)

var objectType = []byte("object")

type Object struct {
	key      []byte
}

func NewObject(key []byte) Node{
	return &Object{
		key:      key,
	}
}

func (m *Object) ToHash(size int) int{
	hashNumber:=int(crc32.ChecksumIEEE(m.key))
	if hashNumber>=0{
		return hashNumber%size
	}
	return (-hashNumber)%size
}

func (m *Object) GetKey() []byte{
	return m.key
}

func (m *Object) SetKey(key []byte){
	m.key = key
}

func (m *Object) GetType() []byte{
	return objectType
}

func (m *Object) Compare(b Node) bool{
	if b == nil{
		return false
	}
	return bytes.Equal(m.key, b.GetKey())
}

func (m *Object) CompareKey(key []byte) bool{
	return bytes.Equal(m.key, key)
}