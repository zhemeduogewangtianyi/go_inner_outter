package _interface

import "testing"

type Programmer interface {
	WriterHelloWorld() string
}

type GoProgrammer struct {
}

func (programmer *GoProgrammer) WriterHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestGoInterface(t *testing.T) {
	var programer Programmer = new(GoProgrammer)
	var result string = programer.WriterHelloWorld()
	t.Log(result)
}
