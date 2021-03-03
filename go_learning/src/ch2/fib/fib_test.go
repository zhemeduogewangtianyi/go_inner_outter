package fib

import (
	"testing"
)

//全局变量
var a int

func TestLibList(t *testing.T) {

	//var a int = 1
	//var b int = 1

	//var(
	//	a int = 1
	//	b = 1
	//)

	//类型推断
	a = 1 //全局变量赋值
	b := 1

	for i := 0; i < 5; i++ {
		temp := a
		a = b
		b = temp + a

		//fmt.Println(a)
		t.Log(a)

	}

}

//交换两个数字
func TestExchange(t *testing.T) {
	var (
		a int = 1
		b int = 2
	)

	//temp := a
	//a = b
	//b = temp

	//一句赋值语句对多个变量赋值
	a, b = b, a

	t.Log(a, b)
}
