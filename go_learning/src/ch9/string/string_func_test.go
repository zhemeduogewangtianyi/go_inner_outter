package string

import (
	"strconv"
	"strings"
	"testing"
)

/**
string 常用函数

	1 : strings
	2 : strconv

*/
func TestStringFn(t *testing.T) {

	s := "A,B,C"

	//分割
	parts := strings.Split(s, ",")
	t.Log(parts)
	printArr(&parts, t)

	s = strings.Join(parts, "-")
	t.Log(s)

}

func printArr(s *[]string, t *testing.T) {
	for index, val := range *s {
		t.Log(index, val)
	}
}

/**
数字转换成字符串
字符串转换成数字
*/
func TestConv(t *testing.T) {

	s := strconv.Itoa(10)
	t.Log(s + " Str ")
	if val, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + val)
	}

}
