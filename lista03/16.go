package main

import "fmt"

func main() {
	var n, k, count int
	fmt.Scanln(&n, &k)

	presentes := make([]int, n)
	var adiantados []int
	for i := 0; i < n; i++ {
		fmt.Scan(&presentes[i])
		if presentes[i] <= 0 {
			count++
			adiantados = append(adiantados, i+1)
		}
	}
	if count >= k {
		fmt.Println("Não")
		for i := 0; i < len(adiantados); i++ {
			fmt.Println(adiantados[i])
		}
	} else {
		fmt.Println("sim")
	}
}
