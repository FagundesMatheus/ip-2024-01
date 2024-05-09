package main

import (
    "fmt"
    "math"
    "sort"
)

func main() {
    var n int
    fmt.Scanln(&n)

    for i := 0; i < n; i++ {
        MenorDistancia()
    }
}

func MenorDistancia() {
    var x, min, comb int
    fmt.Scanln(&x)

    nums := make([]int, x)
    for i := 0; i < x; i++ {
        fmt.Scan(&nums[i])
    }

    sort.Ints(nums)
    min = int(math.Abs(float64(nums[0] - nums[1])))
    for i := 0; i < x-1; i++ {
        if int(math.Abs(float64(nums[i]-nums[i+1]))) < min {
            min = int(math.Abs(float64(nums[i] - nums[i+1])))
        }
    }
    comb = fatorial(x) / (fatorial(x-2) * fatorial(2))
    fmt.Printf("%d %d\n", min, comb)
    fmt.Scanln()
}

func fatorial(num int) int {
    if num == 2 {
        return 2
    }
    return num * fatorial(num-1)
}
