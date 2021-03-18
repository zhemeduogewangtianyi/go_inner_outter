package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

/**
只运行一次 -> 单例模式
*/

//一个单例对象
type Singleton struct {
}

//单例的指针地址
var singleInstance *Singleton

//确保只执行一次
var once sync.Once

//创建单例对象的方式
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

//测试单例对象的在多线程下的创建
func TestGetSingletonObj(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			//输出对象地址 -> 肯定是同一个对象
			fmt.Println(unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
