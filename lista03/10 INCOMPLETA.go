package main

import "fmt"

func main() {
	var n, max int
	fmt.Scanln(&n)
	freq := make([]int, 11)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		freq[nums[i]] += 1
		if nums[i] > max {
			max = i
		}
		if i == n-1 {
			fmt.Printf("nota: %d, %d vezes \n", nums[i], freq[nums[i]])
		}
	}
	for i := 0; i < 11; i++ {

	}
	fmt.Printf("nota: %d, indice %d \n", max, freq[max])
}

//11 5 6 3 4 3 8 7 4 8 6 4
