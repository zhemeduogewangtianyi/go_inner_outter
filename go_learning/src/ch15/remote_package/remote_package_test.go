package remote_package

import "testing"
import cm "../../github.com/easierway/concurrent_map"

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

/**
go 未解决的依赖问题

	1：同一环境下，不同项目使用同一包的不同版本
	2：无法管理对包的特定版本的依赖

	vendor 路径

	随着 Go 1.5 release 版本的发布，vendor 目录被添加到除了 GOPATH 和
	GOROOT 之外的依赖目录查找的解决方案。在 GO 1.6 之前，你需要手动的设置
	环境变量

	查找依赖包路径的解决方案如下：
	1：当前包下的 vendor 目录
	2：向上级目录查找，直到找到 src 下的 vendor 目录
	3：在 GOPATH 下面查找依赖包
	4：在 GOROOT 目录下查找

	常用的依赖管理工具

		godep	https://github.com/tools/godep
		glide	https://github.com/Masterminds/glide
		dep		https://github.com/golang/dep

	cd ch15/remote_package
	glide init
	glide install
*/
