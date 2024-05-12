package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}
	fmt.Print(sum)
}
