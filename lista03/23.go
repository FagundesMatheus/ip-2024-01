package main

import (
	"fmt"
	"math"
)

func main() {
	var str string
	var flag bool = false
	var sum float64

	var A [6]int
	var B [6]int

	fmt.Scan(&str)
	char := []rune(str)

	for i := 0; i < len(char); i++ {
		if string(char[i]) == ";" {
			if flag == false {
				flag = true
			} else {
				fmt.Println("FORMATO INVALIDO!")
				break
			}
		}
		if !flag {
			A[Index(string(char[i]))]++
		}
		if flag {
			B[Index(string(char[i]))]++
		}

	}
	if !flag {
		fmt.Print("FORMATO INVALIDO!")
		return
	}
	for i := 1; i < 6; i++ {
		sum += math.Pow(float64(A[i]-B[i]), 2)
	}
	fmt.Println(math.Sqrt(sum))
}
func Index(str string) int {
	switch str {
	case "a":
		return 1
	case "e":
		return 2
	case "i":
		return 3
	case "o":
		return 4
	case "u":
		return 5
	default:
		return 0
	}
}
