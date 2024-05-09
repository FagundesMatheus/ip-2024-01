package main

import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)

    nums := make([]int, n)
    mapa := make(map[int]bool)

    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
        mapa[nums[i]] = true
    }

    for k := range mapa {
        fmt.Println(k)
    }
}
