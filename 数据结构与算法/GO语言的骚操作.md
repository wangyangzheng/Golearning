#### 变量赋值
```go
    //同时定义两个值
    a, b := 0, 1
    //交换两个值
    a, b = b, a
```
注意使用 `a, b := 1, a + 1` 时后要注意不可以先后影响
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

