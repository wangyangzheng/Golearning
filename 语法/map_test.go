package map_test

import "testing"

//map的声明
func TestMap(t *testing.T) {
	//第一种声明方式
	m1 := map[int]int{1 : 1, 2 : 2}
	t.Log(m1[1])
	//第二种声明方式
	m2 := map[int]int {}
	m2[1] = 1
	t.Log(m2[1])
	//第三种声明方式
	m3 := make(map[int] int, 10)
	m3[1] = 1
	t.Log(m2[1])
}

//测试默认值
func TestMap1(t *testing.T) {
	m2 := map[int]int {1 : 1, 2 : 2}
	m2[3] = 3
	if v, ok := m2[3]; ok{
		t.Logf("Key is  existing %d", v)
	} else {
		t.Log("Key is existing")
	}
}

//遍历整个map
func TestMap2(t *testing.T) {
	m3 := map[int]int {1 : 1, 2 : 2}
	for _, v := range m3 {
		t.Log(v)
	}
}

//前置知识复习
func TestArray(t *testing.T) {
	arr := [...] int {1, 2, 3, 5}
	//idx表示下标,value表示值
	for idx, value := range arr {
		t.Log(idx, value)
	} 
}

func TestArray1(t *testing.T) {
	arr := [...]int {1, 2, 3}
	if ok := arr[0] == 1; ok{
		t.Log(arr[0])
	}
}