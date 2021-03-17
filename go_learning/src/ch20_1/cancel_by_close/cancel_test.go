package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

/**
任务的取消
*/

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func isCancel(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

/**
发一条消息，谁收到关谁
*/
func TestCancel_1(t *testing.T) {
	//创建 channel
	ch := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, ch)
	}

	cancel_1(ch)
	time.Sleep(time.Second * 1)

}

/**
广播发送，全关
*/
func TestCancel_2(t *testing.T) {
	ch := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, ch)
	}
	cancel_2(ch)
	time.Sleep(time.Second * 1)
}
