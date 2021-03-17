package ch18

import (
	"fmt"
	"testing"
	"time"
)

/**
CSP 并发机制
(Communicating sequential processes)
20世纪70年代提出，依赖一个通道来完成两个通讯实体之间的协调

CSP vs. Actor
	1：和 Actor 的直接通讯不同，CSP 模式测试通过 Channel 进行通讯的，更松耦合一些。
	2：Go 中 channel 是有容量限制并且独立于 Groutine，而入 Erlang，Actor 模式中的
	mailbox 容量是无限的，接受进程也总是被动地处理消息

协程会主动的处理从 channel 里面传递过来的信息

Channel 分为两种
1：通讯的两方都同时在线才能通过 channel 进行交互，只要有一方不在，另外一方就会阻塞等待，直到对方完成这个交互
2：buffered channel 消息的接受者和发送者有着更松耦合的关系，通过给 channel 设置容量
*/

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {

	//创建一个 string 通道
	retCh := make(chan string)

	go func() {
		//运行业务代码
		ret := service()
		fmt.Println("returned result.")
		//结果返回到 channel 里面
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	//获取 string channel
	var retCh chan string = AsyncService()
	//运行其他的任务
	otherTask()
	//从 string 通道中取东西
	fmt.Println(<-retCh)
	//time.Sleep(time.Second * 1)
}

/**
buffered channel
*/
func BufferedAsyncService() chan string {

	//声明 channel 的 cap 为 1
	var retCh chan string = make(chan string, 1)

	go func() {
		var ret string = service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh

}

func TestBufferedAsyncService(t *testing.T) {
	var retCh chan string = BufferedAsyncService()
	otherTask()
	var data string = <-retCh
	fmt.Println(data)
}

/**
普通 channel
	working on something else
	returned result.
	service exited.
	Task is done.

	working on something else
	returned result.
	Task is done.
	Done
	service exited.


Buffered Channel
	working on something else
	returned result.
	service exited.
	Task is done.
	Done

区别：
	普通 channel 概率性不等待 service 运行完就结束，意味着普通 channel 会阻塞 service
	Buffered channel 会在 service 执行完之后立马返回结果，不阻塞 service

	Buffered Channel 会更高效一点
*/
