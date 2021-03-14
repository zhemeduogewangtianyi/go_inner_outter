package series

import "fmt"

/**
package

	1：基本复用模块单元
		以首字母大写来表明可被包外代码访问

	2：代码的 package 可以和所在的目录不一致

	3：同一目录里的 Go 代码的 package 要保持一致

*/
func GetFibonacci(n int) []int {

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList

}

/**
方法名小写，包外无法访问
*/
func square(n int) int {
	return n * n
}

func Square(n int) int {
	return n * n
}

/**
init 方法

	1：在 main 被执行前，所有依赖的 package 的 init 方法都会被执行
	2：不同包的 init 函数按照包导入的依赖关系决定执行顺序
	3：每个包可以有多个 init 函数
	4：包的每个源文件也可以有多个 init 函数，这点比较特殊
*/
func init() {
	fmt.Println("my_series.go 的 init 方法被执行")
}

func init() {
	fmt.Println("my_series.go 的第二个 init 方法被执行")
}
