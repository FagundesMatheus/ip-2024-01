package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanln(&N)

	pares := make([]int, 0)
	impares := make([]int, 0)

	var num int
	for i := 0; i < N; i++ {
		fmt.Scanln(&num)
		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
	}

	sort.Ints(pares)
	sort.Sort(sort.Reverse(sort.IntSlice(impares)))

	for _, num := range pares {
		fmt.Println(num)
	}

	for _, num := range impares {
		fmt.Println(num)
	}
}
