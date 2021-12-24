package system

// hex .
type hex struct {
	hexBytes []byte
	binBytes []byte

	hexBytesLen int
	binBytesLen int
}

// NewHex .
func NewHex(data []byte) *hex {
	var hex = new(hex)

	hex.hexBytesLen = len(data)
	hex.binBytesLen = hex.hexBytesLen>>1 + hex.hexBytesLen&1

	hex.hexBytes = data
	hex.binBytes = make([]byte, hex.binBytesLen)

	return hex
}

// Hex2Bin .
func (hex *hex) Hex2Bin() []byte {
	hook(hex.hexBytesLen, hex.conver, hex.split)
	return hex.binBytes
}

// conver 窗口中的十六进制转换成二进制
func (hex *hex) conver(bounds ...int) {
	var startIndex, endIndex = hex.bounds(bounds...)

	// 如果十六进制长度为奇数，则首位单独处理
	if startIndex == 0 && hex.hexBytesLen&1 == 1 {
		hex.binBytes[0] = hex.bin(0, hex.hexBytes[0])
		startIndex++
	}

	for i := startIndex; i < endIndex; i += 2 {
		hex.binBytes[i>>1] = hex.bin(hex.hexBytes[i], hex.hexBytes[i+1])
	}
}

// split .
// 两个4位十六进制组成一个8位二进制,
// 所以需要对左右边界做处理, 确保窗口内为正确的8位二进制
func (hex *hex) split(part int) []*window {
	var (
		windows = make([]*window, part)

		// 此为转换为二进制时，所计算出得窗口边界
		// 实际使用时需转换为十六进制下标
		x, y = hex.binBytesLen / part, hex.binBytesLen % part
	)

	for i := range windows {
		switch i {
		case 0:
			windows[i] = &window{
				l: 0,
				// 如果十六进制长度为奇数，则第一位为低四位
				r: x<<1 + hex.hexBytesLen&1,
			}
		case part:
			windows[i] = &window{
				l: windows[i-1].r,
				r: hex.hexBytesLen,
			}
		default:
			if y > 0 {
				windows[i] = &window{
					l: windows[i-1].r,
					r: windows[i-1].r + (x+1)<<1,
				}
				y--
			} else {
				windows[i] = &window{
					l: windows[i-1].r,
					r: windows[i-1].r + x<<1,
				}
			}
		}
	}

	return windows
}

// bound .
func (hex *hex) bounds(v ...int) (int, int) {
	switch len(v) {
	case 0:
		return 0, hex.hexBytesLen
	case 1:
		return v[0], hex.hexBytesLen
	default:
		return v[0], v[1]
	}
}

// bin .
func (hex *hex) bin(higt, low byte) byte {
	return higt<<4 | low
}
