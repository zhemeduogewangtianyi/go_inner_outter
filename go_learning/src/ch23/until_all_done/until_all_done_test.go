package until_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
所有任务都完成
*/
func runTask(id int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("The result is from %d \r\n", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)

	for i := 0; i < 10; i++ {
		go func(id int) {
			ch <- runTask(id)
		}(i)
	}

	finalRet := ""

	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\r\n"
	}

	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(1 * time.Second)
	t.Log("After", runtime.NumGoroutine())
}
