package polymorphic

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriterHelloWorld() Code
}

type GoProgrammer struct {
}

func (goProgrammer *GoProgrammer) WriterHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (javaProgrammer *JavaProgrammer) WriterHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

func printPolyMorPhic(p Programmer) {
	fmt.Printf("%T %v \r\n", p, p.WriterHelloWorld())
}

func printPolyMorPhicPointer(p *Programmer) {
	fmt.Printf("%T %v \r\n", (*p), (*p).WriterHelloWorld())
}

func TestPolyMorPhic(t *testing.T) {

	javaProgrammer := new(JavaProgrammer)
	goProgrammer := new(GoProgrammer)
	printPolyMorPhic(javaProgrammer)
	printPolyMorPhic(goProgrammer)

	//???????
	//javaProgrammerPointer := new(JavaProgrammer)
	//goProgrammerPointer := new(GoProgrammer)
	//printPolyMorPhicPointer(&javaProgrammerPointer)
	//printPolyMorPhicPointer(&goProgrammerPointer)

	java := &JavaProgrammer{}
	goLang := GoProgrammer{}
	printPolyMorPhic(java)
	printPolyMorPhic(&goLang)
}
