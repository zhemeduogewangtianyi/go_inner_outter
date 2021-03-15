package remote_package

import "testing"
import cm "github.com/easierway/concurrent_map"

/**
package

	1：通过 go get 来获取远程依赖
		go get -u 强制从网络更新远程依赖
	2：注意代码在 github 上的组织形式，以适应 go get
		直接以代码路径开始，不要有 src

	示例：
		https://github.com/easierway/concurrent_map
		go get -u github.com/easierway/concurrent_map
*/
func TestConcurrentMap(t *testing.T) {

	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))

}
