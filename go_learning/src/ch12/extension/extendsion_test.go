package extension

import (
	"fmt"
	"testing"
)

/**
定义 Pet 结构体
定义 speak 方法
定义 speakTo (name string) 方法
*/
type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...PET...")
}

func (p *Pet) speakTo(content string) {
	p.speak()
	fmt.Println(" ", content)
}

func (p *Pet) wocao() {
	fmt.Println("握草")
}

/**
1: 复合 Pet
2: 为 Dog 定义与 Pet 同样的方法
*/
type Dog struct {
	p *Pet
	Pet
}

func (dog *Dog) speak() {
	fmt.Println("dog......")
}

func (dog *Dog) speakTo(content string) {
	dog.speak()
	fmt.Println(" ", content)
}

func TestExtension(t *testing.T) {

	/**
	var dog Pet := new(Dog)
	不符合 LSP 规则 ， 因为 GO 不支持显示类型转换
	不支持重载
	*/
	dog := new(Dog)
	dog.speakTo("dogggggg")
	dog.p.wocao()
	dog.wocao()
}
