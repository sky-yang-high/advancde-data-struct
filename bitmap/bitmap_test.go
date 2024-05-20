package bitmap_test

import (
	"advanced-data-struct/bitmap"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := bitmap.NewBitMap(10)
	bm.Set(10)

	if bm.Get(10) != true {
		t.Error("bitmap set/get test failed")
	}
	if bm.Get(11) != false {
		t.Error("bitmap get test failed")
	}
	if bm.Count() != 1 {
		t.Error("bitmap set/count test failed")
	}

	bm.Clear(10)
	if bm.Get(10) != false {
		t.Error("bitmap clear/get test failed")
	}
}
