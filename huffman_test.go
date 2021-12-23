package go_huffman

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	{
		var start = time.Now()
		for i := 0; i < 1<<20; i++ {
			strconv.Itoa(i)
		}
		fmt.Println("1 -->", time.Since(start))
	}
	{
		var start = time.Now()
		var s sync.WaitGroup
		s.Add(8)
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		go func() {
			for i := 0; i < 131072; i++ {
				strconv.Itoa(i)
			}
			s.Done()
		}()
		s.Wait()
		fmt.Println("2 -->", time.Since(start))
	}
}

func TestSlice(t *testing.T) {
	fmt.Println(5>>1 + 5&1)
}
