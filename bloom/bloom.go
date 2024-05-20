package bloom

import (
	"advanced-data-struct/bitmap"
	"fmt"
)

type BloomFilter struct {
	bitmap    bitmap.Bitmap
	k         uint
	m         uint //the bits of bitmap, same as size*8 of bitmap
	n         uint //same as count of bitmap
	encryptor *Encryptor
}

func New(size, k uint, encryptor *Encryptor) *BloomFilter {
	return &BloomFilter{
		bitmap:    *bitmap.NewBitMap(size),
		m:         size * 8,
		k:         k,
		encryptor: encryptor,
	}
}

func (bf *BloomFilter) Add(data string) {
	bf.n++
	ec := bf.getEncryptedHash(data)
	for _, h := range ec {
		bf.bitmap.Set(h)
	}
}

func (bf *BloomFilter) getEncryptedHash(data string) []uint {
	ec := make([]uint, bf.k)
	for i := 0; i < int(bf.k); i++ {
		ecd := uint(bf.encryptor.Encrypt([]byte(data)))
		ec[i] = ecd % uint(bf.m)
		data = fmt.Sprint(ecd)
	}
	return ec
}

func (bf *BloomFilter) Exists(val string) bool {
	ec := bf.getEncryptedHash(val)
	for _, h := range ec {
		if !bf.bitmap.Get(h) {
			return false
		}
	}
	return true
}
