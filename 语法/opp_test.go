package opp_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
	
}

func TestCreatEmployeeObj(t *testing.T) {
	e := Employee {"0", "Bob", 20}
    e1 := Employee {Name : "Tom", Age : 30}
    e2 := new(Employee)
    e2.Id = "2"
    e2.Age = 22
    e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	//输出类型e2的类型
	t.Logf("%T", e2)
}


// func (e Employee) String() string {
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

func (e *Employee) String() string {
	//输出一下传进来的地址
	fmt.Printf("Adress is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
func TestStruct(t *testing.T) {
	//e := Employee{"0", "Bob", 20}
	e := &Employee{"0", "Bob", 20}
	//此时e是一个指针,指针调用的时候不需要箭头
	//Adress is c0000b6760
	fmt.Printf("Adress is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
