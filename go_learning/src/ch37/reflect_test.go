package ch37

import (
	"fmt"
	"reflect"
	"testing"
)

/**
reflect.TypeOf vs reflect.ValueOf

	reflect.TypeOf 返回类型 reflect.Type

	reflect.ValueOf 返回值 reflect.Value

	可以从 reflect.Value 获得类型

	通过 kind 的来判断类型
*/
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
	//t.Log(reflect.ValueOf(f).Addr())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)

	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

/**
利用反射编写灵活的代码

按名字访问结构的成员
reflect.ValueOf(*e).FieldByName("Name")

按名字访问结构的方法
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
*/
type Employee struct {
	Employee string

	/**
	Struct Tag

	type BasicInfo struct {
		Name string `format:"normal"`
		Age int `json:"age"`
	}
	*/

	Name string `format:"normal"`
	Age  int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func (c *Customer) UpdateAge(newVal int) {
	c.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name : value(%[1]v) , Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age : ", e)

}
