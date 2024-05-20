package bloom

import (
	"github.com/spaolacci/murmur3"
)

type Encryptor struct {
}

func NewEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) Encrypt(data []byte) uint32 {
	hasher := murmur3.New32()
	hasher.Write(data)
	return hasher.Sum32()
}
