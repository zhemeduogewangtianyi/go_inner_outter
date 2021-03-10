package empty_interface

import (
	"fmt"
	"testing"
)

/**
go 语言中的任意类型
c 中的 void*
java 中的 Object

*/

func DoSomething(p interface{}) {

	/**
	两段式 if else 语句
	p.(int) 代表断言
	*/

	if i, ok := p.(int); ok {
		fmt.Printf("%[1]T %[1]d \r\n", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Printf("%[1]T %[1]s \r\n", s)
		return
	}

	fmt.Printf("Unkown Type %[1]T \r\n", p)

}

func DoSomethingSwitch(p interface{}) {
	switch val := p.(type) {
	case int:
		fmt.Printf("%[1]T %[1]d \r\n", val)
	case string:
		fmt.Printf("%[1]T %[1]s \r\n", val)
	default:
		fmt.Printf("Unkown Type %[1]T \r\n", p)
	}
}

type Dog struct {
}

/**
go 接口最佳实践

1：倾向于使用小的接口，很多接口只包含一个方法 --- single method interface

	type Reader interface{
		Reader(p byte[]) (n int , err error)
	}

	type Writer interface{
		Writer(p byte[]) (n int , err error)
	}

2：较大的接口的定义，可以由多个小借口定义组合而成

	type ReaderWriter interface{
		Reader
		Writer
	}

3：只依赖于必要功能的最小接口

	func StoreData (reader Reader) error {
		...
	}

*/
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(1)
	DoSomething("1")
	DoSomething(new(Dog))

	DoSomethingSwitch("ppp")
	DoSomethingSwitch(-10)
	DoSomethingSwitch(new(Dog))

}
