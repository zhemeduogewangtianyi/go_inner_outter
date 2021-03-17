package share_mem

import (
	"sync"
	"testing"
	"time"
)

/**
共享内存并发机制

	package sync

		Mutex
		RwLock
*/
func TestCounter(t *testing.T) {

	var counter int = 0
	for i := 0; i < 5000; i++ {
		go func(num int) {
			counter++
		}(i)
	}
	time.Sleep(1 * time.Second * 7)
	t.Logf("counter = %d \r\n", counter)

}

/**
var mut = sync.Mutex
defer (){
	mut.Unlock()
}()
mut.Lock()
*/
func TestCounterThreadSafe(t *testing.T) {

	var counter int = 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	//不睡眠的话会导致主线程完毕，子线程未完成
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)

}

/**
WaitGroup 类似于 java 的 java.util.concurrent.CountDownLatch 的 await() + countDown() 但是又不是 countDownLatch
*/
func TestCounterWaitGroup(t *testing.T) {

	var counter int = 0
	var lock sync.Mutex
	var countDownLatch sync.WaitGroup

	countDownLatch.Add(5000)

	for i := 0; i < 5000; i++ {

		//countDownLatch.Add(1)

		go func() {
			defer func() {
				countDownLatch.Done()
				lock.Unlock()
			}()
			lock.Lock()
			counter++
		}()
	}
	countDownLatch.Wait()
	t.Logf("counter %d", counter)

}
