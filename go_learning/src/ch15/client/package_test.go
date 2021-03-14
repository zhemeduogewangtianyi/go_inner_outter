package client

import "testing"
import x "../../ch15/series"

/*
	去调用 my_series.go 的斐波那契数列
*/

func TestPackage(t *testing.T) {
	fibonacci := x.GetFibonacci(5)
	t.Log(fibonacci)
	//t.Log(series.square(5))
	t.Log(x.Square(5))
}
