package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	s = "hello"
	//测试长度
	t.Log(len(s))

	//表示一个字叫做严
	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	t.Log(len(s)) //中的长度是3

	c := []rune(s)
	//观察中的unicode值
	t.Logf("中 unicode %x", c[0])
	t.Log("中 UTF-8", s)
}

func TestArray(t *testing.T) {
	//string 是一个不可以改变的切片
	var str string
	str = "hello"
	t.Log(str)
	//string 是一个不可以改变的切片
	//这是错误写法str[1] = '3'
}

func TestString1(t *testing.T) {
	var str string
	str = "中"
	c := []rune(str)
	t.Logf("%x", str)  //e4 b8 ad
	t.Logf("%x", c[0]) //4e2d
}

//字符串的遍历
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国家"
	for _, c := range s {
		//这里的格式化表示的是都与第一个对应
		t.Logf("%[1]c %[1]d", c)
	}
}

//字符串的相关方法
func TestFunction(t *testing.T) {
	//字符串分割
	s := "a,b,c"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	//字符串的连接
	t.Log(strings.Join(parts, "-"))
	//字符串和整数的转化
	s = strconv.Itoa(10)
	t.Log("str" + s)
	//直接使用会发生错误
	//t.Log(10 + strconv.Atoi("10"))
	if i , err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
