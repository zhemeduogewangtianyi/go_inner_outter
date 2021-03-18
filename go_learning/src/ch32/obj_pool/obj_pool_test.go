package obj_pool

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

/**
对象池

GOLAND 不能运行，只能切换到文件目录下运行一下命令
什么原因。。。
go test -v .\obj_pool_test.go .\obj_pool.go
*/
func TestObjPool(t *testing.T) {

	pool := NewObjPool(10)

	for i := 0; i < 11; i++ {
		if reusable, error := pool.GetObj(time.Millisecond * 100); error != nil {
			t.Error(error)
		} else {
			t.Log(reusable)
			t.Logf("%T \r\n", reusable)
			t.Log(unsafe.Pointer(reusable))
			if error := pool.ReturnObj(reusable); error != nil {
				t.Error(error.Error())
			} else {
				fmt.Println("return Obj to Pool success !")
			}

		}
	}
}
