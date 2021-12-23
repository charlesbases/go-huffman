package go_huffman

type window struct {
	l int
	r int
}

type bytes interface {
	// conver 转化窗口中的进制
	// if len(bounds) == 0, 转换全部字节
	// if len(bounds) == 1, 转换指定位置到结尾的全部字节
	// if len(bounds) == 2, 转化窗口中的进制
	conver(bounds ...int)
	// Result 转换结果
	result() []byte

	// len 原始数据长度
	len() int
	// split 数据长度等分
	split(part int) []*window
}

// bin2hex 二进制转十六进制
func bin2hex(v []byte) []byte {
	return NewPromise(NewBin(v)).Result()
}

// hex2bin 十六进制转二进制
func hex2bin(v []byte) []byte {
	return NewPromise(NewHex(v)).Result()
}

// bin 二进制
type bin struct {
	ori []byte
	res []byte

	orilen int
	reslen int
}

// NewBin .
func NewBin(data []byte) bytes {
	var bin = new(bin)

	bin.ori = data
	bin.orilen = len(data)
	bin.reslen = bin.orilen << 1
	bin.res = make([]byte, bin.reslen)

	return bin
}

// conver 窗口中的二进制转换成十六进制
func (bin *bin) conver(bounds ...int) {
	var oriStartIndex, oriEndIndex = bin.bounds(bounds...)
	for i := oriStartIndex; i < oriEndIndex; i++ {
		bin.res[i<<1], bin.res[i<<1+1] = bin.hex((bin.ori)[i])
	}
}

// result .
func (bin *bin) result() []byte {
	return bin.res
}

// len .
func (bin *bin) len() int {
	return bin.orilen
}

// split 等分
// eg: split(10, 3) ==> [0, 3), [3, 7), [7, 10)
//     split(9, 3)  ==> [0, 3), [3, 6), [6, 9)
func (bin *bin) split(part int) []*window {
	var (
		windows = make([]*window, part)
		x, y    = bin.orilen / part, bin.orilen % part
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
				r: bin.orilen,
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

// bounds .
func (bin *bin) bounds(v ...int) (int, int) {
	switch len(v) {
	case 0:
		return 0, bin.orilen
	case 1:
		return v[0], bin.orilen
	default:
		return v[0], v[1]
	}
}

// hex 二进制转化成十六进制
func (bin *bin) hex(v byte) (byte, byte) {
	return v >> 4, v & 0xf
}

// hex 十六进制
type hex struct {
	ori []byte
	res []byte

	orilen int
	reslen int
}

// NewHex .
func NewHex(data []byte) bytes {
	var hex = new(hex)

	hex.ori = data
	hex.orilen = len(data)
	hex.reslen = hex.orilen>>1 + hex.orilen&1
	hex.res = make([]byte, hex.reslen)

	return hex
}

// conver 窗口中的十六进制转换成二进制
func (hex *hex) conver(bounds ...int) {
	var oriStartIndex, oriEndIndex = hex.bounds(bounds...)

	// 如果十六进制长度为奇数，则首位单独处理
	if oriStartIndex == 0 && hex.orilen&1 == 1 {
		hex.res[0] = hex.bin(0, hex.ori[0])
		oriStartIndex++
	}

	for i := oriStartIndex; i < oriEndIndex; i += 2 {
		hex.res[i>>1] = hex.bin(hex.ori[i], hex.ori[i+1])
	}
}

// result .
func (hex *hex) result() []byte {
	return hex.res
}

// split .
// 两个4位十六进制组成一个8位二进制,
// 所以需要对左右边界做处理, 确保窗口内为正确的8位二进制
func (hex *hex) split(part int) []*window {
	var (
		windows = make([]*window, part)

		// 此为转换为二进制时，所计算出得窗口边界
		// 实际使用时需转换为十六进制下标
		x, y = hex.reslen / part, hex.reslen % part
	)

	for i := range windows {
		switch i {
		case 0:
			windows[i] = &window{
				l: 0,
				// 如果十六进制长度为奇数，则第一位为低四位
				r: x<<1 + hex.orilen&1,
			}
		case part:
			windows[i] = &window{
				l: windows[i-1].r,
				r: hex.orilen,
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

// len .
func (hex *hex) len() int {
	return hex.orilen
}

// bounds .
func (hex *hex) bounds(v ...int) (int, int) {
	switch len(v) {
	case 0:
		return 0, hex.orilen
	case 1:
		return v[0], hex.orilen
	default:
		return v[0], v[1]
	}
}

// bin .
func (hex *hex) bin(higt, low byte) byte {
	return higt<<4 | low
}
