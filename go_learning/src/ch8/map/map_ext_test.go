package _map

import "testing"

/**
map 的扩展
*/
func TestMapWithFunValue(t *testing.T) {

	/**
	map 中可以存方法。
	*/
	m := map[int]func(op int) int{}

	var f1 = func(op int) int {
		return op
	}
	m[1] = f1

	var f2 = func(op int) int {
		return op * op
	}
	m[2] = f2

	var f3 = func(op int) int {
		return op * op * op
	}
	m[3] = f3

	for k, v := range m {
		t.Log(k, v(2))
	}

}

/**
go 语言中没有 Set，可以用 map 自己实现一个 map[type]bool

增加
判断有没有
删除
元素个数
*/
func TestMapForSet(t *testing.T) {

	m := map[int]bool{}

	//增
	m[1] = true

	//元素个数
	t.Log(len(m))

	//判断元素有没有
	t.Log(conditionElementExists(&m, 1))

	//删
	delete(m, 1)

	t.Log(conditionElementExists(&m, 1))

}

func conditionElementExists(m *map[int]bool, element int) bool {

	//if val , ok := m[element] ; ok {
	//	return val
	//}else{
	//	return false
	//}

	if val, ok := (*m)[element]; ok {
		return val
	} else {
		return false
	}

}
