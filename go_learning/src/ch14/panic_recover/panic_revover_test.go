package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/*

panic recover

	panic

		panic 用于不可恢复的错误
		panic 退出前会执行 defer 指定的内容

	panic vs os.Exit

		os.Exit 退出时不会调用 defer 指定的函数
		os.Exit 退出时不输出当前调用栈信息
*/
func TestPanicVsExit(t *testing.T) {

	fmt.Println("Start")

	os.Exit(-1)

}

func TestExit(t *testing.T) {

	fmt.Println("Start")

	os.Exit(-1)

}

/*
	会输出调用栈信息
*/
func TestPanic(t *testing.T) {

	fmt.Println("Start")
	panic(errors.New("Something wrong !"))

}

/**
测试 defer 在 panic 之后会执行
*/
func TestPanicAndDefer(t *testing.T) {
	defer func() {
		fmt.Println("Finally !")
	}()

	fmt.Println("Start")

	panic(errors.New("Something wrong !"))

}

/**
recover

	在 panic 之后会调用 defer 关键字标识的函数 ， defer 函数里面可以通过 recover() 获取到 panic 传进去的 error

	类似于 java try-catch-finally 中的 catch 块

	defer func () {
		if err := recover() ; err != nil {
			err something
		}
	}()
*/
func TestRecover(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from ", err)
		}
	}()

	fmt.Println("Start")

	panic(errors.New("Something Wrong !"))

}
