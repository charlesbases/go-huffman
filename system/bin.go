package system

// bin 二进制
type bin struct {
	binBytes []byte
	hexBytes []byte

	binBytesLen int
	hexBytesLen int
}

// NewBin .
func NewBin(data []byte) *bin {
	var bin = new(bin)

	bin.binBytesLen = len(data)
	bin.hexBytesLen = bin.binBytesLen << 1

	bin.binBytes = data
	bin.hexBytes = make([]byte, bin.hexBytesLen)

	return bin
}

// Bin2Hex .
func (bin *bin) Bin2Hex() []byte {
	hook(bin.binBytesLen, bin.conver, bin.split)
	return bin.hexBytes
}

// conver 窗口中的二进制转换成十六进制
func (bin *bin) conver(bounds ...int) {
	var startIndex, endIndex = bin.bound(bounds...)
	for i := startIndex; i < endIndex; i++ {
		bin.hexBytes[i<<1], bin.hexBytes[i<<1+1] = bin.hex((bin.binBytes)[i])
	}
}

// split 等分
// eg: split(10, 3) ==> [0, 3), [3, 7), [7, 10)
//     split(9, 3)  ==> [0, 3), [3, 6), [6, 9)
func (bin *bin) split(part int) []*window {
	var (
		windows = make([]*window, part)
		x, y    = bin.binBytesLen / part, bin.binBytesLen % part
	)

	for i := range windows {
		switch i {
		case 0:
			windows[i] = &window{
				l: 0,
				r: x,
			}
		case part:
			windows[i] = &window{
				l: windows[i-1].r,
				r: bin.binBytesLen,
			}
		default:
			if y > 0 {
				windows[i] = &window{
					l: windows[i-1].r,
					r: windows[i-1].r + x + 1,
				}
				y--
			} else {
				windows[i] = &window{
					l: windows[i-1].r,
					r: windows[i-1].r + x,
				}
			}
		}
	}

	return windows
}

// bound .
func (bin *bin) bound(v ...int) (int, int) {
	switch len(v) {
	case 0:
		return 0, bin.binBytesLen
	case 1:
		return v[0], bin.binBytesLen
	default:
		return v[0], v[1]
	}
}

// hex .
func (bin *bin) hex(v byte) (byte, byte) {
	return v >> 4, v & 0xf
}
