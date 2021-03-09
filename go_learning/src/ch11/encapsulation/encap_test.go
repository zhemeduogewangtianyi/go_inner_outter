package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

//c struct
/**
struct Student {
	char name[30] ;
	unsigned int age ;



} Stu , *stu;
*/

type Employee struct {
	id   string
	name string
	age  int
}

/**
三种创建结构体的方式 struct

	1 :
		e := {"id","name", 3}
	2 :
		e := {id : "id" , name : "name" , age : 1}
	3 :
		e := new (Employee)
		e.id = "id"
		e.name = "name"
		e.age = 1
*/
func TestCreateEmployeeObj(t *testing.T) {

	e := Employee{"0", "laoye", 18}
	t.Log(e)

	e1 := Employee{id: "1", name: "yeye", age: 18}
	t.Log(e1)

	e2 := new(Employee)
	e2.id = "2"
	e2.name = "baba"
	e2.age = 18
	t.Log(e2, e2.name)

	t.Logf("e is %T \r\n   e1 is %T \r\n  e2 is %T \r\n", e, e1, e2)

	t.Logf(" e change to e2 type : %T", &e)

	t.Log(e.String())
	t.Log(e.String1())

	t.Log(e1.String())
	t.Log(e1.String1())

	t.Log(e2.String())
	t.Log(e2.String1())

}

/**
为 Employee 结构体定义一个方法
*/
func (e Employee) String() string {
	//通过 unsafe 查看指针地址
	fmt.Printf("Address is %x \r\n", unsafe.Pointer(&e.name))
	return fmt.Sprintf("id:%s-name:%s-age:%d", e.id, e.name, e.age)
}

func (e *Employee) String1() string {
	//通过 unsafe 查看指针地址
	fmt.Printf("Address is %x \r\n", unsafe.Pointer(&e.name))
	return fmt.Sprintf("id:%s-name:%s-age:%d", e.id, e.name, e.age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"1", "2", 10}

	t.Logf("e Address is : %x", e)

	t.Log(e.String())
	//指针可以直接访问，不用 (&e)->String1()
	t.Log((&e).String1())
}
