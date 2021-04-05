package ch38

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/**
万能程序：

	DeepEqual : 比较切片和 map
*/
func TestDeppEqual(t *testing.T) {

	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}
	//t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))

}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

type Employee struct {
	Name string
	Age  int
}

/**
关于“反射”你应该知道的
• 提⾼了程序的灵活性
• 降低了程序的可读性
• 降低了程序的性能
*/
func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 10}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	//func (v Value) Elem() Value
	//Elem returns the value that the interface v contains or that the pointer
	//It panics if v's Kind is not Interface or Ptr.
	//It returns the zero Value if v is nil

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		//Elem() 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil

}
