package slice_test

import (
	"strconv"
	"testing"
)

/**
创建切片
*/
func TestSliceInit(t *testing.T) {

	//第一种声明式创建方式
	var s0 []int
	t.Log(len(s0), cap(s0))

	//填充元素
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	//第二种类型推断方式
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	//3: 通过 make( type , len ,cap) 创建
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	/*
		len 代表可访问元素的个数
		cap 代表容量
	*/
	//t.Log(s2[0] , s2[1] , s2[2] , s2[3] , s2[4])
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])

	//证明切片是可变长的
	s2 = append(s2, 3)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5])
	t.Log(len(s2), cap(s2))
}

/**
测试切片可变长
*/
func TestSliceGrowing(t *testing.T) {

	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	/**

		=== RUN   TestSliceGrowing
	    slice_test.go:52: 1 1
	    slice_test.go:52: 2 2
	    slice_test.go:52: 3 4
	    slice_test.go:52: 4 4
	    slice_test.go:52: 5 8
	    slice_test.go:52: 6 8
	    slice_test.go:52: 7 8
	    slice_test.go:52: 8 8
	    slice_test.go:52: 9 16
	    slice_test.go:52: 10 16

		1：当切片的 cap（容量） 不够的时候，append( type , data) 会触发 2倍 扩容
		2：为什么要 s = append( type , data )，是因为 append 方法在触发扩容的时候会重新开辟一块连续的数组，将老数组的元素 copy 进新数组，
		内存地址会发生变化。
		3：切片的这种做法是有代价的，性能调优的时候会说

		优点：
			不用像数组那样预估元素个数，切片会自增长
		缺点：
			更多存储空间的复制是有代价的

	*/

}

/**
切片的共享内存
*/
func TestSliceShareMemory(t *testing.T) {

	var year []string
	for i := 0; i < 12; i++ {
		//year[i] = string(i) + "月"
		year = append(year, strconv.Itoa(i+1)+"月")
		//year[i] = strconv.Itoa(i + 1) + "月"
	}

	t.Log(year, len(year), cap(year))

	spring := year[0:3]
	t.Log(spring, len(spring), cap(spring))

	summer := year[4:7]
	t.Log(summer, len(spring), cap(spring))

	//修改切片的数据
	summer[0] = "null"

	//共享内存 year 的元素也会发生变化
	t.Log(year, len(spring), cap(spring))
	t.Log(spring, len(spring), cap(spring))
	t.Log(summer, len(spring), cap(spring))

}

/**
测试切片是否可以比较
不能比较
*/
func TestSliceCompare(t *testing.T) {

	/**
	数组和切片的区别
	1：容量是否可伸缩
	2：是否可以进行比较
		数组的比较：两个维度相同的数组，比较数组里面的每一个元素是否都相同
		例如：arr1 == arr2
	*/

	//var s1 []int = []int{1,2,3,4}
	//var s2 []int = []int{1,2,3,4}

	//Invalid operation: s1 == s2 (operator == is not defined on []int)
	//t.Log(s1 == s2)

}
