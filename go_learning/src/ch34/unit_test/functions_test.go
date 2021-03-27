package unit

import (
	"fmt"
	"testing"
)

/**
go test -v .\functions.go .\functions_test.go

单元测试
*/
func TestSquare(t *testing.T) {
	inpouts := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inpouts); i++ {
		ret := square(inpouts[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inpouts[i], expected[i], ret)
		}
	}
	t.Log(square(10))
}

/**
内置单元测试框架：

Fail,Error：该测试失败，该测试据需，其他测试继续执行
FailNow, Fatal：该测试失败，该测试终止，其他测试继续执行

代码覆盖率
	go test -v -cover
断言
	http://github.com/stretchr/testify

	go get -u github.com/stretchr/testify/assert
*/
func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error !")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error !")
	fmt.Println("End")
}

func TestSquareAssert(t *testing.T) {
	inpouts := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inpouts); i++ {
		ret := square(inpouts[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inpouts[i], expected[i], ret)
		}
	}
	t.Log(square(10))
}
