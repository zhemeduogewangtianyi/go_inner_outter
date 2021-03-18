package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
sync.Pool 对象缓存

一、sync.Pool 对象获取

	1：尝试从私有对象获取 - 私有对象是协程安全的，共享池是协程不安全的
	2：私有对象不存在，尝试从当前 Processor 的共享池获取
	3：如果当前 Processor 共享池也是空的，那么就尝试去其他 Processor 的共享池获取
	4：如果所有子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回

二、sync.Pool 对象放回

	1：如果私有对象不存在则保存为私有对象
	2：如果私有对象存在，放入当前 Processor 子池的共享池中

三、使用 sync.Pool

	pool := &sync.Pool {
		New : func() interface{} {
			return 0
		},
	}

	arr := pool.Get().(int)
	pool.put(10)

四、sync.Pool 对象的生命周期

	1：GC 会清除 sync.Pool 缓存的对象
	2：对象的缓存有效期为下一次 GC 之前
		由于 GC 是系统级别调度，无法清除的知道一个对象的生命周期。怎么长时间控制一个连接的生命周期，完全做不到？

*/
func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)

	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolGc(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	if val, ok := pool.Get().(int); ok {
		fmt.Println(val)
	}

	pool.Put(3)
	if val1, ok := pool.Get().(int); ok {
		fmt.Println(val1)
	}

	//拿过了就没了。。
	if val1, ok := pool.Get().(int); ok {
		fmt.Println(val1)
	}

	pool.Put(4)
	if val2, ok := pool.Get().(int); ok {
		fmt.Println(val2)
	}

	pool.Put(5)

	//主动调用 gc
	runtime.GC()

	if val2, ok := pool.Get().(int); ok {
		fmt.Println(val2)
	}

}

/**
多协程 sync.Pool 是如何工作
*/
func TestSyncPoolMultiGoroutine(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			obj := pool.Get().(int)
			t.Log(obj)
			wg.Done()
		}(i)
	}
	wg.Wait()

	/**
	sync.Pool 总结

		1：适合于通过复用，降低复杂对象的创建和 GC 的代价
		2：协程安全，会有锁的开销
		3：生命周期受 GC 影响，不适合于做连接池等，需要自己管理生命周期的资源的池化
	*/

}
