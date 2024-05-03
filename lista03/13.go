package main

import "fmt"

func main() {
	var n, x, Fmax, num int
	fmt.Scanln(&n)

	freq := make([]int, 101)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		freq[x+1] += 1
	}
	for i := 0; i < 101; i++ {
		if freq[i] > Fmax {
			Fmax = freq[i]
			num = i - 1
		}
	}
	fmt.Printf("%d\n%d", num, Fmax)
}
