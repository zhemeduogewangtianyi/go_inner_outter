package error

import (
	"errors"
	"fmt"
	"testing"
)

/**
普通的获取费布那切数列
*/
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

/**
go 多返回结果
异常机制
*/
func GetFibonacciNew(n int) ([]int, error) {

	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2 , 100]")
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

/**
自定义异常
*/
var LessThanTwoError = errors.New("n should be less than 2")
var LargerThanHundredError = errors.New("n should be larger than 100")

/**
区分不同的异常
*/
func GetFibonacciErrorType(n int) ([]int, error) {

	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {

	t.Log(GetFibonacci(10))

	t.Log("\r\n")

	if array, error := GetFibonacciNew(1); error != nil {
		t.Error(error.Error())
	} else {
		t.Log(array)
	}

	t.Log("\r\n")

	if arr, err := GetFibonacciErrorType(1); err != nil {
		t.Error(err)
	} else {
		t.Log(arr)
	}

	if arr, err := GetFibonacciErrorType(101); err != nil {
		if err == LessThanTwoError {
			fmt.Println("小了")
		}
		if err == LargerThanHundredError {
			fmt.Println("大了")
		}
		t.Error(err)
	} else {
		t.Log(arr)
	}

	t.Log("\r\n")

	if arr, err := GetFibonacciErrorType(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("小了")
		}
		if err == LargerThanHundredError {
			fmt.Println("大了")
		}
		t.Error(err)
	} else {
		t.Log(arr)
	}
}
