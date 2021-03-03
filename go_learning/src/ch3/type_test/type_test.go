package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {

	var a int32 = 1
	var b int64 = 1

	//cannot use a (type int) as type int64 in assignment
	//b = a
	b = int64(a)

	var c MyInt
	//c = b
	c = MyInt(b)

	t.Log(a, b, c)

}

func TestPoint(t *testing.T) {

	var a int = 1
	var aPtr *int = &a

	arr := [3]int{1, 2, 3}

	t.Log(arr)

	/**
	go 语言不支持指针移动
	*/
	//t.Log(a,aPtr,aPtr + 1)

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

}

func TestString(t *testing.T) {

	var s string

	t.Logf("*%s*", s)
	t.Log(len(s))

	if s == "" {
		t.Log("字符串是空串")
	}

}
