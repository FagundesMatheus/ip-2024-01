package main

import "fmt"

func main() {
	var testes int
	fmt.Scanln(&testes)

	for i := 0; i < testes; i++ {
		validar()
	}
}
func validar() {
	var b1, b2 int
	nums := make([]int, 11)
	for i := 0; i < 11; i++ {
		fmt.Scan(&nums[i])
	}
	for i := 0; i < 9; i++ {
		b1 += nums[i] * (i + 1)
		b2 += nums[i] * (9 - i + 1)
	}
	if b1%11 == 10 {
		b1 = 0
	} else {
		b1 = b1 % 11
	}
	if b2%11 == 10 {
		b2 = 0
	} else {
		b2 = b2 % 11
	}
	if nums[9] == b1 && nums[10] == b2 {
		fmt.Println("válido")
	} else {
		fmt.Println("inválido")
	}
}

//7 3 3 1 8 4 6 8 0 9 6
//go run lista03\18.go
