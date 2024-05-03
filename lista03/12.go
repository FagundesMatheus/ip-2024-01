package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", nums[i])
	}
}
