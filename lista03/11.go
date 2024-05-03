package main

import "fmt"

func main() {
	var n, max, min int
	fmt.Scanln(&n)

	nums := make([]int, n)
	numsI := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if i == 0 {
			min = nums[0]
		}
		numsI[n-i-1] = nums[i]
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	fmt.Printf("%d\n %d\n%d\n%d", nums, numsI, max, min)
}
