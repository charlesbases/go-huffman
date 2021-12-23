package go_huffman

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

// loop 数据字节长度
var (
	bytesize = 1 * GiB
	bindata  []byte
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestBytes(t *testing.T) {
	_random()

	_hex := bin2hex(bindata)

	_bin := hex2bin(_hex)

	fmt.Println(_verify(bindata, _bin))
}

func _random() {
	bindata = make([]byte, bytesize)
	for i := range bindata {
		bindata[i] = uint8(rand.Intn(1<<8 - 1))
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
