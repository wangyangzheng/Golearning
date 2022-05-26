package loop_test

import (
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 1
	//while (n < 5) {}
	for n < 5 {
		t.Log(n)
		n++
	}
}

//condition
func TestCondition(t *testing.T) {
	n := 1
	if n > 1 {
		t.Log("数字大于1")
	} else if n < 1 {
		t.Log("数字小于1")
	} else {
		t.Log("数字等于1")
	}
}

//第二种写法
func TestCondition2(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("a为1")
	}
}
//第二种写法在函数的返回值中用的很多

// func TestCodition3(t *testing.T) {
// 	// error表示函数返回的错误, sum为一个函数
// 	if a, error := sum(); error == nil{
// 		t.Log("1 == 1")
// 	} else {
// 		t.Log("error")
// 	}
// }

func TestCodition4(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not ")
		}		
	}
}

