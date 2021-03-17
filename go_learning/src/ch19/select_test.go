package ch19

import (
	"fmt"
	"testing"
	"time"
)

/**
多路选择和超时控制

1：多渠道选择
	不能依赖 case 的顺序
	select {
		case ret := <-retCh1 :
			t.Logf("result %s \r\n" , ret)
		case ret := <-retCh2 :
			t.Logf("result %s \r\n" , ret)
		default :
			t.Error("No one returned")
	}

2：超时控制

	select {
		case ret := <-retCh :
			t.Logf("result %s \r\n" , ret)
		case <-time.After(time.Second * 1) :
			t.Error("time out")
	}
*/

func service() string {
	//time.Sleep(time.Millisecond * 50)
	time.Sleep(time.Millisecond * 500)
	return "Service Done"
}

func AsyncService() chan string {
	//创建一个 buffered channel
	var retCh chan string = make(chan string, 1)

	go func() {
		var ret string = service()
		fmt.Println("returned result.")
		//把 service 的结果放进 channel
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func TestSelect(t *testing.T) {

	select {

	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")

	}
}

func WeiXinService() string {
	time.Sleep(time.Millisecond * 100)
	return "WEI_XIN"
}

func QQService() string {
	time.Sleep(time.Millisecond * 100)
	return "Q_Q"
}
func DingDingService() string {
	time.Sleep(time.Millisecond * 300)
	return "D_D"
}

func MoreChannelService() []chan string {

	retWeiXin := make(chan string)
	retQQ := make(chan string)
	retDD := make(chan string)

	var resArr []chan string = []chan string{retWeiXin, retQQ, retDD}

	go func() {
		retWeiXin <- WeiXinService()
	}()

	go func() {
		retQQ <- QQService()
	}()

	go func() {
		retDD <- DingDingService()
	}()

	return resArr

}

func TestMoreChannelService(t *testing.T) {
	retArr := MoreChannelService()
	for i := 0; i < len(retArr); i++ {
		t.Log(<-retArr[i])
	}
}

func WeiXinAsyncService() chan string {

	retCh := make(chan string, 1)

	go func() {
		fmt.Println("WeiXinAsyncService run")
		retCh <- WeiXinService()
	}()

	return retCh
}

func QQAsyncService() chan string {

	retCh := make(chan string)

	go func() {
		fmt.Println("QQAsyncService run")
		retCh <- QQService()
	}()

	return retCh
}

func DingDingAsyncService() chan string {

	retCh := make(chan string)

	go func() {
		fmt.Println("DingDingAsyncService run")
		retCh <- DingDingService()
	}()

	return retCh
}

/**
如果有多个 case 都可以运行，Select 会公平地选出一个执行，其他的不会执行
否则：
	如果有 default 子句，则执行该语句
	如果没有 default 子句，select 将阻塞，知道某个通讯可以运行；Go 不会重新对 channel 或值进行求值

	所有的case处于阻塞状态，就会走default分支
*/
func TestAsyncSelect(t *testing.T) {

	select {

	case ret := <-WeiXinAsyncService():
		t.Log(ret)
	case ret := <-QQAsyncService():
		t.Log(ret)
	case ret := <-DingDingAsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 200):
		t.Error(" time out")
		//default:
		//	t.Error("Not Found AsyncService")

	}
}
