package unsafe_programming

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/**
"不安全"行为的危险性

i:=10
f := *(*float64)(unsafe.Pointer(&i))
*/
func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	//这种类型转换很危险  5e-323
	t.Log(f)
}

type MyInt int

//合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 10)
			}
			wg.Done()
		}()
	}

	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			readDataFn()
			time.Sleep(time.Microsecond * 100)
		}
		wg.Done()
	}()

	wg.Wait()
}
