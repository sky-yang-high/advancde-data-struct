package bloom_test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.T) {
	var a uint32 = 4294967295
	var b int32 = int32(a)
	fmt.Println(b)
}
