package custom_type

import (
	"fmt"
	"testing"
	"time"
)

// ch10 -> func_test.go 的计时模式，装饰者模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(o int) int {
		start := time.Now()
		ret := inner(o)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}

func lowFun(op int) int {
	time.Sleep(5 * time.Second)
	fmt.Println("卧槽")
	return op
}

/**
自定义类型
*/
type IntConv = func(op int) int

func timeSpentNew(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}

func TestCustomType(t *testing.T) {
	var intConv IntConv = timeSpentNew(lowFun)
	t.Log(intConv(5))
}
