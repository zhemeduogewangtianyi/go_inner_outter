package ch20_2

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
Context 与任务取消

1：Context

	1：根 Context : 通过 context.Background() 创建
	2：子 Context : 通过 context.WithCancel( parentContext ) 创建
		2.1：	ctx , context = context.WithCancel( context.Background() )
	3：当前 Context 被取消时，基于他的子 context 都会被取消
	4：接受取消通知 <-ctx.Done()
*/
func isCancel(cxt context.Context) bool {
	select {
	case <-cxt.Done():
		return true
	default:
		return false
	}
}

func TestPreClassLookUp(t *testing.T) {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	createGoFunc(ctx1, 5, "ch1")
	createGoFunc(ctx2, 5, "ch2")
	cancel1()
	cancel2()
	time.Sleep(time.Second * 3)
}

func createGoFunc(ctx context.Context, len int, name string) {
	for i := 0; i < len; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancel(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(name, i, "Done")
		}(i, ctx)
	}
}
