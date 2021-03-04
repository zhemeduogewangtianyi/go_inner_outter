package _map

import (
	"strconv"
	"testing"
)

/**
初始化 Map
*/
func TestInitMap(t *testing.T) {

	m1 := map[string]int{
		"One":   1,
		"Tow":   2,
		"Three": 3,
	}
	t.Log(m1, len(m1))

	m2 := map[string]string{}
	m2["key"] = "value"
	t.Log(m2, len(m2))

	/**
	1 : make 创建 map 的时候 make( map , cap) 不能 make( map , size , cap)
	2 : map 不能用 cap(map) 访问 cap ...
	*/
	m3 := make(map[string]string, 16)
	m3["k"] = "v"
	t.Log(m3, len(m3))

}

/**
从 map 里面取元素的时候，map 是不是空？ key 在 map 中是否存在？ key 所对应的 value 是否 null？
*/
func TestAccessNotExsitingKey(t *testing.T) {

	var m map[string]string

	t.Log("未初始化", m["0"])

	m = map[string]string{}

	//取不到会有默认值
	t.Log("初始化没元素", m["0"])

	m["3"] = "000010"

	//判断 map 获取到的数据是否为空  if 要判断的变量 , 返回值(boolean) ; 返回值(boolean) {  xxxx } else { yyyy }
	if val, ok := m["3"]; ok {
		t.Log("Key 3 value is ", val)
	} else {
		t.Log("Key 3 is not existing")
	}

}

/**
测试 map 的遍历
*/
func TestTravelMap(t *testing.T) {
	m := map[string]string{}
	for i := 0; i < 5; i++ {
		m[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	for k, v := range m {
		t.Log(k, v)
	}

	for k, _ := range m {
		t.Log(k)
	}

	for _, v := range m {
		t.Log(v)
	}

}
