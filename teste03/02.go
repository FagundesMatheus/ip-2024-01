package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var len int
		fmt.Scanln(&len)
		nums := make([]int, len)
		for i := 0; i < len; i++ {
			fmt.Scan(&nums[i])
		}
		fmt.Println(TrocaOpostoSeMenor(len, nums))
	}
}
func troca(len int, nums []int, i int) []int {
	nums[i], nums[len-1-i] = nums[len-1-i], nums[i]
	return nums
}
func TrocaOpostoSeMenor(len int, nums []int) []int {
	for i := 0; i < len; i++ {
		x := nums[i]
		y := nums[len-1-i]
		if i < len-1-i && x < y {
			nums = troca(len, nums, i)
		}
	}
	return nums
}
