package dunc

import (
	"fmt"
	"math/rand"
	"strings"
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

func UserLoginInterceptor(inner func(password string) bool) func(password string) bool {

	return func(pwd string) bool {
		if ok := strings.EqualFold(pwd, "123456"); ok {
			pwd = "654321"
			return inner(pwd)
		} else {
			return false
		}
	}

}

func UserLogin(password string) bool {
	time.Sleep(time.Second * 3)
	fmt.Println(password + " 登录成功")
	return true
}

func TestUserLogin(t *testing.T) {

	var method func(password string) bool = UserLoginInterceptor(UserLogin)
	var result bool = method("12345")
	t.Log(result)

	t.Log(UserLoginInterceptor(UserLogin)("123456"))

}

/**
可变长参数
*/
func Sum(ops ...int) int {
	var ret int
	for _, val := range ops {
		ret += val
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

/**
defer 关键字

	类似于 try catch finally 中的 finally

panic 关键字

	类似于 Exception异常
*/
func TestDefer(t *testing.T) {

	defer func() {
		fmt.Println("defer 1")
	}()

	fmt.Println("1111111111111111111111111")

	defer TempDeferFn()

	panic("java.lang.Exception")

}

func TempDeferFn() {
	fmt.Println("defer 2")
}
