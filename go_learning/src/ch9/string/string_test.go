package string

import "testing"

func TestString(t *testing.T) {

	var s string

	//string 初始化默认 "" , size = 0
	printString(s, t)

	s = "hello"
	printString(s, t)

	//string 是不可变的 byte slice   -> can not assign to s[1]
	//s[1] = '3'

	//可以存储任何 2进制 数据
	s = "\xE4\xB8\xA5"
	//s = "\xE4\xBa\xB5\xFF"
	printString(s, t)

	/**
	unicode 是一种字符集（code point）
	UTF8 是 unicode 的存储实现（转换为字节序列的规则）

	rune 是一个类型，可以帮我们取出字符串的 unicode, go 语言的内置机制，自动转换。属于切片
	*/

	s = "中"
	//此时的 len(s) 输出的是 byte 数
	printString(s, t)

	var c []rune = []rune(s)
	t.Logf("%s unicode %x , len : %d", s, c[0], len(c))
	t.Logf("%s    UTF8 %x , len : %d", s, s, len(s))
}

func printString(s string, t *testing.T) {
	t.Logf("s is : %s    *****    len is : %d", s, len(s))
}

/**
字符串遍历
*/
func TestStringToRune(t *testing.T) {
	s := "大疯狂老虎机aB哈1acd"

	for index, val := range s {
		//%[1] 代表对应第一个参数 c:字符 d:数字 x:16进制编码
		t.Logf("%[1]c %[1]d %d %[1]x", val, index)
	}
}
