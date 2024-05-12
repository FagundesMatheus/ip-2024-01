
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var max float64
	fmt.Scanln(&n)
	x := make([]int, n)
	y := make([]int, n)
	z := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&x[i], &y[i], &z[i])
		if i > 0 {
			if math.Abs(float64(x[i]-x[i-1])) > max {
				max = math.Abs(float64(x[i] - x[i-1]))
			}
			if math.Abs(float64(y[i]-y[i-1])) > max {
				max = math.Abs(float64(y[i] - y[i-1]))
			}
			if math.Abs(float64(z[i]-z[i-1])) > max {
				max = math.Abs(float64(z[i] - z[i-1]))
			}
			fmt.Printf("%.2f \n", max)
			max = 0
		}

	}
}
