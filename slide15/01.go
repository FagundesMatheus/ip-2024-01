package main

import "fmt"

func main() {
	var arr [2]int
	fmt.Scanln(&arr[0], &arr[1])
	fmt.Print(recursiva(arr[0], arr[1]))
}

func recursiva(num int, expoente int) int {
	if expoente == 1 {
		return num
	}
	return num * recursiva(num, expoente-1)
}
