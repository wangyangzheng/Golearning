package type_test

import (
	"math"
	"testing"
)

type MyInt int64

//不支持隐式的转换
func TestImplicit1(t *testing.T) {
	var a int64 = 1
	var b int32
	b = int32(a)
	var c MyInt
	c = MyInt(a)
	t.Log(a, b, c)
}

//预定义的值
func TestDefine(t *testing.T) {
	t.Log(math.MaxInt64)
	t.Log(math.MaxFloat64)
	t.Log(math.MaxUint32)
}

//指针不支持加减
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(&a, aPtr)
	//输出类型
	t.Logf("%T %T", a, aPtr)
}

//字符串默认是空的
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
