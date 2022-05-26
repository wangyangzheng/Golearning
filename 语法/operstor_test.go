package operstor_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)//flase
	t.Log(a == d)//true
}

func TestOperator1(t *testing.T) {
	//a := [...] int {1, 2, 3, 4}
	//b := [...] int {1, 2, 3, 4, 5}
	//t.Log(a == b)
	//出现错误invalid operation: a == b (mismatched types [4]int and [5]int)
	t.Log(1)
}

func TestOperator2(t *testing.T) {
	//0111
	a := 7
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
func TestOperator3(t *testing.T) {
	a := 1//01
	t.Log(a &^ 3)//3先变成00然后与1与
}
