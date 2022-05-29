```go
package main


import "fmt"

var (
    n int
    path = [10] int{}
    st = [10] bool{}
)


func dfs(u int) {
    if u == n {
        for i := 0; i < n; i++ {
            fmt.Printf("%d ", path[i])
        }
        fmt.Println()
    }
    for i := 1; i <= n; i++ {
        if st[i] == false {
            path[u] = i
            st[i] = true
            dfs(u + 1)
            st[i] = false
        }
    }
}


func main(){
   
    fmt.Scan(&n)
    dfs(0)
}

```
