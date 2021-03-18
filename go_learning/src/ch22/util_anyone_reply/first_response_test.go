package util_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
假设从多个数据源获取数据，只要有一个 csp 协程取到了数据，立马返回结果
*/
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10

	//ch := make(chan string)

	/**
	使用 buffered channel 解决协程阻塞导致泄漏的问题
	*/
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {

	//输出当前系统中的协程数量
	t.Log("Before", runtime.NumGoroutine())

	t.Log(FirstResponse())

	time.Sleep(time.Second * 1)

	t.Log("After", runtime.NumGoroutine())

}
