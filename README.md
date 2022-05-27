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
m3 := make(map[int] int, 10)
```
