package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	//这里的value变成了一个方法的名字
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](3))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	//判断某个元素是否存在
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is existing", n)
	}
	//输出数组的长度
	mySet[3] = true
	t.Log(len(mySet))

	//删除一个元素
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
