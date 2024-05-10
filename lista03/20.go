/*package main

import "fmt"

func main() {
	var n, max2 int
	n1 := make([]int, 3)
	n2 := make([]int, 3)
	fmt.Scanln(&n)
	for j := 1; j < n; j++ {

		if j == 1 {
			for i := 0; i < 2; i++ {
				fmt.Scan(&n1[i])
			}
		} else {
			for i := 0; i < 3; i++ {
				n1[i] = n2[i]
			}
		}
		for i := 0; i < 2; i++ {
			fmt.Scan(&n2[i])
			if n2[i]-n1[i] > max2 {
				max2 = n2[i] - n1[i]
			}
		}
		fmt.Println(max2)
	}

} */

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
