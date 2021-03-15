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

	for i := 0; i < 5000; i++ {
		countDownLatch.Add(1)
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

/**
CSP 并发机制
(Communicating sequential processes)
19世纪70年代提出，依赖一个通道来完成两个通讯实体之间的协调

CSP vs. Actor
	1：和 Actor 的直接通讯不同，CSP 模式测试通过 Channel 进行通讯的，更松耦合一些。
	2：Go 中 channel 是有容量限制并且独立于 Groutine，而入 Erlang，Actor 模式中的
	mailbox 容量是无限的，接受进程也总是被动地处理消息

协程会主动的处理从 channel 里面传递过来的信息

Channel 分为两种
1：通讯的两方都同时在线才能通过 channel 进行交互，只要有一方不在，另外一方就会阻塞等待，直到对方完成这个交互
2：buffered channel 消息的接受者和发送者有着更松耦合的关系，通过给 channel 设置容量
*/
