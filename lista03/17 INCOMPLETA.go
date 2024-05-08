package main

import "fmt"

func main() {
	var n, count int
	var mapa map[int]int
	fmt.Scanln(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if mapa[nums[i]] == 0 {
			mapa[nums[i]]++
			count++
		}
	}
	fmt.Println(count)

}
