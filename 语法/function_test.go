package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)	
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b) 
	
}

//计算函数操作的时长
//输入是函数类型,返回也是函数类型
func timeFn(inner func (op int) int) func (op int) int {
	return func (n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

//传入的参数
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op 
}

func TestFn1(t *testing.T) {
	tsSF := timeFn(slowFun)
	t.Log(tsSF(10))
}