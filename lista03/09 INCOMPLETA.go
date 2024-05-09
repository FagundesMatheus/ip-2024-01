package main

import (
    "fmt"
    "math"
)

func main() {
    var N int
    fmt.Scanln(&N)

    var x, y, z float64
    var prevX, prevY, prevZ float64
    fmt.Scanln(&prevX, &prevY, &prevZ)

    for i := 1; i < N; i++ {
        fmt.Scanln(&x, &y, &z)
        dist := math.Sqrt(math.Pow(x-prevX, 2) + math.Pow(y-prevY, 2) + math.Pow(z-prevZ, 2))
        fmt.Printf("%.2f\n", dist)
        prevX, prevY, prevZ = x, y, z
    }
}
