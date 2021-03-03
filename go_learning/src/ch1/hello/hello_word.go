package main //包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
) //引入代码依赖

//功能实现
//func main() int {
func main() {

	/**

	编译运行：go run hello_world.go

	1:必须是 main 包 : package main
	2:必须是 main 方法 : func main()
	3:文件名不一定是 main.go

	*/

	fmt.Println("Hello World")

	/**

	Go 中 main 函数不支持任何返回值
	通过 os.Exit 来返回状态

	*/

	//return 1

	//os.Exit(-1)
	//os.Exit(0)

	/**
	main 函数不支持传入参数
	func main(args []String) 错误

	在程序中直接通过 os.Args 获取命令行参数 -> go run ./hello_world.go www.baidu.com
	*/

	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("os.Args 有参数：", os.Args[1])
	}

}
