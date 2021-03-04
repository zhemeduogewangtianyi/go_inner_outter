package dunc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
函数
1：可以有多个返回值
2：所有参数都是值传递：slice，map，channel 会有传引用的错觉
3：函数可以作为变量的值
4：函数可以作为参数和返回值
*/
func returnReturnMutiValues() (int, int, int) {
	return rand.Intn(10), rand.Intn(20), rand.Intn(30)
}

/**
记录时间
输入函数 ，返回函数
*/
func timeSpent(inner func(op int) int) func(op int) int {

	return func(n int) int {

		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent : ", time.Since(start).Seconds())
		return ret
	}

}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

/**
测试函数
*/
func TestFn(t *testing.T) {
	a, b, _ := returnReturnMutiValues()
	t.Log(a, b)
	tsSf := timeSpent(slowFun)
	var result int = tsSf(10)
	t.Log(result)
}
