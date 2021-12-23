package go_huffman

import (
	"runtime"
	"sync"
)

var DefauleLines = runtime.NumCPU()

func init() {
	runtime.GOMAXPROCS(DefauleLines)
}

// Promise .
type Promise struct {
	base bytes
}

// NewPromise .
func NewPromise(data bytes) *Promise {
	var pro = new(Promise)
	pro.base = data

	switch {
	case pro.base.len() == 0:
	// nothing
	case pro.base.len() > 1<<20:
		var wg sync.WaitGroup

		for _, item := range pro.base.split(DefauleLines) {
			wg.Add(1)

			go func(w *window) {
				pro.base.conver(w.l, w.r)
				wg.Done()
			}(item)
		}

		wg.Wait()
	default:
		pro.base.conver()
	}

	return pro
}

// Result .
func (pro *Promise) Result() []byte {
	return pro.base.result()
}
