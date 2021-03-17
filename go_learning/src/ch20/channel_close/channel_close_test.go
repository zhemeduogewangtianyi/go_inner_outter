package channel_close

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

/**
channel 的关闭

1：向关闭的 channel 发送数据，会导致 panic

2：v , ok <-ch ; ok 为 boolean 值，true 表示正常接受，false 表示通道关闭

3：所有的 channel 接受者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 ok 值为 false。
这个广播机制常被利用，进行向多个订阅者同事发送信号。
例如：退出信号
*/

func dataProducer(ch chan int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

		//关闭的 channel 继续发消息会受到一个 panic
		//ch<-11

		wg.Done()
	}()

}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i := 0 ; i < 10 ; i++ {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel closed. data : " + strconv.Itoa(data))
				break
			}

		}
		wg.Done()
	}()
}

/**
不判断 <-ch 继续从关闭的 channel 里面取值，会取到 chan int 返回的 0
*/
func noConditionDataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 11; i++ {

			var data int = <-ch
			fmt.Println("noConditionDataReceiver : " + strconv.Itoa(data))

		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)

	//多个 receiver
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)

	//不判断
	wg.Add(1)
	noConditionDataReceiver(ch, &wg)

	wg.Wait()
}
