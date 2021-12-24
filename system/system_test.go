package system

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	Byte int = 1
	KiB      = 1 << 10 * Byte
	MiB      = 1 << 10 * KiB
	GiB      = 1 << 10 * MiB
)

var (
	size     = 1 * GiB
	binBytes []byte
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBin_Hex(t *testing.T) {
	_random()

	_hex := NewBin(binBytes).Bin2Hex()

	_bin := NewHex(_hex).Hex2Bin()

	fmt.Println(_verify(binBytes, _bin))
}

func _random() {
	binBytes = make([]byte, size)
	for i := range binBytes {
		binBytes[i] = uint8(rand.Intn(1<<8 - 1))
	}
}

// _verify .
func _verify(v1, v2 []byte) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true
}
