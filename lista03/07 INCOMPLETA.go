package main

import "fmt"

func main() {
    var n, max int
    for {
        fmt.Scanln(&n)
        if n == 0 {
            break
        }

        nums := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scan(&nums[i])
            if nums[i] > max {
                max = nums[i]
            }
        }

        m := make(map[int]int)
        for _, num := range nums {
            m[num]++
        }

        for i := 0; i <= max; i++ {
            fmt.Printf("(%d) %d\n", i, m[i])
        }
    }
}
