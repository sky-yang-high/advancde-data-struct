package bitmap

type Bitmap struct {
	bits  []byte //  a byte slice to store the bitmap
	count uint   //  the number of bits in the bitmap
	size  uint   //  the size of the bitmap in bytes
}

func NewBitMap(size uint) *Bitmap {
	return &Bitmap{
		size: size,
		bits: make([]byte, size), //a additional byte
	}
}

// ignore the expansion of the bitmap
func (b *Bitmap) Set(num uint) {
	if num >= b.size*8 {
		return
	}
	b.count++
	b.bits[num/8] |= 1 << (num % 8)
}

func (b *Bitmap) Get(num uint) bool {
	if num >= b.size*8 {
		return false
	}
	return b.bits[num/8]&(1<<(num%8)) != 0
}

func (b *Bitmap) Clear(num uint) {
	if num >= b.size*8 {
		return
	}
	b.count--
	//&^= , means & and ^, in this case, it means clear the bit
	b.bits[num/8] &^= 1 << (num % 8)
}

func (b *Bitmap) Count() uint {
	return b.count
}
