package main

import "fmt"

func main() {
    var n, count int
    fmt.Scanln(&n)

    nums := make([]int, n)
    freq := make(map[int]int)

    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
        freq[nums[i]]++
    }

    for _, v := range freq {
        if v == 1 {
            count++
        }
    }

    fmt.Println(count)
}
