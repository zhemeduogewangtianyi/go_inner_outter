package _interface

import "testing"

/**
接口和实现类
*/
type Programmer interface {
	WriterHelloWorld() string
}

/**
实现类
*/
type GoProgrammer struct {
}

/**
绑定方法 DuckType -> 方法签名和返回值一样

*/
func (goProgrammer *GoProgrammer) WriterHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	var str string = p.WriterHelloWorld()
	t.Log(str)
}
