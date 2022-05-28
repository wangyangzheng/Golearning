# Golearning

#### 学习路径

由郝林哥整理
![lPAVZ.jpg](https://s1.328888.xyz/2022/05/26/lPAVZ.jpg)


#### 第一个程序
1. Go语言的main函数没有返回值，但是可以使用`os.Exit(0)`进行退出

#### 变量和数据类型

1. Go语言不支持变量的隐式类型的转化,但是支持显示的类型转换
2. Go语言支持数据类型的推导
3. Go语言的bool类型只有两个值分别是`true`和`false`
4. Go语言的常量可以自动赋值
5. Go语言的指针类型不可以直接进行加加减减
6. string类型默认是空字符串


#### 运算符
1. Go语言将前缀的++和--删除了
2. Go语言支持`&^`这个符号的含义是先将右边的数字进行取反，然后在与左边的数进行`&`运算


#### 循环语句和循环语句
1. Go语言中支持一边赋值,一边判断条件的条件语句
2. Go语言只支持`for`循环
3. Go语言的`switch`不需要使用`break`这个关键字

```go
func TestArray1(t *testing.T) {
	arr := [...]int {1, 2, 3}
	if ok := arr[0] == 1; ok{
		t.Log(arr[0])
	}
}
```

```go
func TestArray(t *testing.T) {
	arr := [...] int {1, 2, 3, 5}
	//idx表示下标,value表示值
	for idx, value := range arr {
		t.Log(idx, value)
	} 
}
```

#### 数组
1. 数组的声明
```go
var a int [3]
arr := [3] int {1, 2, 3}
//二维数组
arr := [2][2] int {{1, 2}, {3, 4}}
```
2. 数组是可以进行截取的
```golang
a[1, 2]
```
3. 数组的长度
可以使用`len()`函数
#### 切片
1. 切片有长度，个数， 容量
2. 切片的内存是共享的

#### map
1. map有两部分组成一部分是`键`,一部分是`值`
2. map的三种声明方式
```go
m1 := map[int]int{1 : 1, 2 : 2}
m2 := map[int]int {}
//这里的10指的是容量
m3 := make(map[int] int, 10)
```
3. 如果map种不存在某个键,那么系统自动为它的值赋上0

#### 字符串
1. 字符串的切片不可以改变
2. 字符串的编码是unicode,编码规则是utf-8
3. 字符串的相关操作

#### 函数
1. 函数可以有多个返回值
2. 函数的返回值可以是函数
3. 函数的参数是可变长的


关于函数可以有多个返回值的演示
```go
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)	
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b) 
	
}
```

函数的返回值也可以是函数
```go
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

```


关于可变长参数的演示
```go
//opd是可变长参数
func Sum(ops ...int) int {
	res := 0
	for _, op := range ops {
		res += op
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}
```
