package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var mediana float32
	fmt.Scanln(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	if n%2 == 1 {
		mediana = float32(nums[n/2])
	} else {
		mediana = float32(nums[n/2])/2 + float32(nums[n/2-1])/2
	}
	fmt.Printf("%.2f", mediana)
}

//0 1 2 3 4 5
