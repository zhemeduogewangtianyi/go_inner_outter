package series

/**
package

	1：基本复用模块单元
		以首字母大写来表明可被包外代码访问

	2：代码的 package 可以和所在的目录不一致

	3：同一目录里的 Go 代码的 package 要保持一致

*/
func GetFibonacci(n int) []int {
	fiboList := []int{1, 1}

	for i := 2; i < n; i++ {
		fiboList = append(fiboList, fiboList[i-2]+fiboList[i-1])
	}
	return fiboList
}
