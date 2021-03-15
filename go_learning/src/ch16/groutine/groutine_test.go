package groutine

import (
	"fmt"
	"testing"
	"time"
)

/**
Thread vs. Groutine (线程和协程)

1：创建时默认 stack 的大小
	JDK5 以后 java Thread stack 默认为 1M
	Groutine 的 stack 初始化大小为 2k

2：和 KSE (Kernel Space Entity) 的对应关系
	Java Thread 是 1:1
	Groutine 是 M:N
*/

/**
测试协程
*/
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {

		//会把 i 复制一份单独用作 num 作为参数，不会有多线程对同一份资源竞争的情况发生
		go func(num int) {
			fmt.Println(num)
		}(i)

		//会并发，会抢占同一地址的 i
		//go func(){
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 50)
}
