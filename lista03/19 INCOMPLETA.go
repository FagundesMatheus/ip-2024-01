package main

import "fmt"

func main() {
	var n int
	var mapa map[int]int
	fmt.Scanln(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if mapa[nums[i]] == 0 {
			mapa[nums[i]]++
		}
	}
	fmt.Println(mapa) //tentar printar apenas as chaves do map

}
