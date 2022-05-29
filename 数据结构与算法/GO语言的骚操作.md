#### 变量赋值
```go
    //同时定义两个值
    a, b := 0, 1
    //交换两个值
    a, b = b, a
```
注意: 假如我们想让一个数更新完之后去更新另一个数字，写成这样就会出错
```go
package main

import "fmt"


func main(){
    a, b := 1, 1
    //假如我们想让a先更新然后更新b, 如果这样写就会出错
    a, b = a + 1, a + 1
    fmt.Println(a, b)//输出
}
```
#### 循环语句
```go
    //idx表示数组的下标, value表示数组的值
    for idx, value := range nums {
        fmt.Print(idx, value)
    }
```
省略idx
```go
for _, value := range nums {
        fmt.Println(value)
    }
```
用for循环模仿do while
```go
    for {
        i++
        if condition {
            break
        }
    }
```
```
```
#### 切片的定义
```go
num := make([]int, 1000000)
```

