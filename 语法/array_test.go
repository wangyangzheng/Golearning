package arry_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	//arr2 := [4] int {1, 2, 3, 4}
	//arr3 := [...] int {1, 2, 3, 4}
	t.Log(arr[1], arr[2])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	//第二种遍历方式
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
	//第二种遍历方式改版写法可以使用下滑线表示这玩意可以省略
	for _, e := range arr3 {
		t.Log(e)
	}
}

//多维数组的使用

func TestArray2(t *testing.T) {
	//表示这是一个二维数组
	b := [2][2]int{{1, 2}, {3, 4}}
	//二维数组的遍历
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			t.Log(b[i][j])
		}
	}
}

// 数组的截取
func TestArray3(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	t.Log(a[1:2]) //表示2
	b := a[:]
	t.Log(b)
}

//这是切片
func TestSlice(t *testing.T) {
	var s0 []int
	//len()是数组的长度,cap表示数组的容量
	t.Log(len(s0), cap(s0))

	//填入一个元素
	s0 = append(s0, 1)

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	//表示长度是3,容量是5
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	//这样会发生越界t.Log(s2[0], s2[1], s2[2], s3[3])
	s2 = append(s2, 1)
	t.Log(s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		//结构体的内存进行扩展
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", " aug", "sep", "Qct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	//看看这两个数组的地址
	summer[0] = "Unknow"
	t.Log(Q2)
}
