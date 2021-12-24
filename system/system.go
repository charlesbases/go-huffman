package system

import (
	"runtime"
	"sync"
)

var (
	// defaultLinie 并行限制
	defaultLinie = runtime.NumCPU()
	// defaultSizeLimit 字节大小限制，超过则启用并行转换
	defaultSizeLimit = 1 << 20
)

func init() {
	runtime.GOMAXPROCS(defaultLinie)
}

type window struct {
	l int
	r int
}

// hook .
func hook(size int, conver func(bounds ...int), split func(part int) []*window) {
	switch {
	case size == 0:
		// nothing
	case size > defaultSizeLimit:
		var swg sync.WaitGroup

		for _, item := range split(defaultLinie) {
			swg.Add(1)

			go func(w *window) {
				conver(w.l, w.r)

				swg.Done()
			}(item)
		}

		swg.Wait()
	default:
		conver()
	}
}
