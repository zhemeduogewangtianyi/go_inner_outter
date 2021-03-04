package array_test

import "testing"

/**
创建数组
*/
func TestArrayInit(t *testing.T) {

	//声明数组
	var arr [3]int
	arr[1] = 5
	t.Log(arr[1], arr[2])

	//声明并且初始化
	arr1 := [3]int{1, 2, 3}
	t.Log(arr1[0], arr1[1])

	arr2 := [...]int{1, 2, 3}
	t.Log(arr2[0], arr2[1])

}

/**
数组遍历
*/
func TestArrayTravel(t *testing.T) {

	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	/**
	 c : 	int size = sizeof(arr) / sizeof(arr[0]);
	java/js : 	int size = arr.length;
	go :	int size = len(arr)
	*/
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	arr1 := [...]int{100, 200, 300, 400, 500, 600, 700, 800, 900}

	/**
	range 遍历数组：
		index 代表索引下标
		value 代表值
	*/
	for index, value := range arr1 {
		t.Log(index, value)
	}

	/**
	不关心的值可以用 _ 占位，保证编译通过
	*/
	for _, value := range arr1 {
		t.Log(value)
	}

}

/**
二维数组遍历
*/
func TestMutidimensionalArray(t *testing.T) {

	var mutidimensionalArray [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i := 0; i < len(mutidimensionalArray); i++ {
		for j := 0; j < len(mutidimensionalArray[i]); j++ {
			t.Log(mutidimensionalArray[i][j])
		}
	}

	for _, value := range mutidimensionalArray {
		for _, val := range value {
			t.Log(val)
		}
	}

}

/**
数组截取
*/
func TestArraySecsion(t *testing.T) {

	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	arr1 := arr[3:]
	t.Log(arr1)

	arr2 := arr[:3]
	t.Log(arr2)

	arr3 := arr[:]
	t.Log(arr3)

	arr4 := arr[0:len(arr)]
	t.Log(arr4)

}
