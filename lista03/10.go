package main

import "fmt"

func main() {
    var n, max, maxIndex int
    fmt.Scanln(&n)
    freq := make([]int, 11)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
        freq[nums[i]]++
        if nums[i] > max {
            max = nums[i]
            maxIndex = i
        }
    }
    fmt.Printf("nota: %d, %d vezes\n", nums[n-1], freq[nums[n-1]])
    fmt.Printf("nota: %d, indice %d\n", max, maxIndex)
}
