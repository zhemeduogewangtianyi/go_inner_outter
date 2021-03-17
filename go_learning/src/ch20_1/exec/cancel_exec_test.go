package exec

import (
	"fmt"
	"testing"
	"time"
)

/**
练习
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

func TestPreClassLookUp(t *testing.T) {
	ch1 := make(chan struct{}, 0)
	ch2 := make(chan struct{}, 0)
	createGoFunc(ch1, 5, "ch1")
	createGoFunc(ch2, 5, "ch2")
	cancel_1(ch1)
	cancel_2(ch2)
	time.Sleep(time.Second * 1)
}

func createGoFunc(ch chan struct{}, len int, name string) {
	for i := 0; i < len; i++ {
		go func(i int, ch chan struct{}) {
			for {
				if isCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(name, i, "Done")
		}(i, ch)
	}
}
