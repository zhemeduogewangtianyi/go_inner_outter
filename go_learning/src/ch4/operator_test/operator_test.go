package operator_test

import "testing"

/**
运算符
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}

	//数组比较
	t.Log(a == b)

	//长度不同会报出长度不同的错误，编译不通过
	//t.Log(a == c)

	t.Log(a == d)
}

/**
&^ 按位清零
*/
func TestBinarySetZero(t *testing.T) {

	t.Log(1 &^ 1)
	t.Log(2 &^ 1)
	t.Log(3 &^ 1)
	t.Log(4 &^ 1)

	t.Log(1 &^ 0)
	t.Log(2 &^ 0)
	t.Log(3 &^ 0)
	t.Log(6 &^ 0)

}

const (
	READER = 1 << iota
	WRITER
	EXECUTE
)

/**
bit 位清除
*/
func TestBitClear(t *testing.T) {
	//0111
	var a int = 7
	a = a &^ 3
	t.Log(a)
	t.Log(READER, WRITER, EXECUTE)
	t.Log(a&READER == READER, a&WRITER == WRITER, a&EXECUTE == EXECUTE)
}
