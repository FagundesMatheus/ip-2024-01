package main

import (
	"fmt"
)

func main() {
	for {
		var n, d int
		fmt.Scanln(&n, &d)

		if n == 0 && d == 0 {
			break
		}

		var number string
		fmt.Scanln(&number)

		max := CalcMax(number, d)
		fmt.Println(max)
	}
}

func CalcMax(number string, d int) string {
	var stack []byte

	for i := 0; i < len(number); i++ {
		digit := number[i]

		for len(stack) > 0 && stack[len(stack)-1] < digit && len(stack)+len(number)-i > d {
			stack = stack[:len(stack)-1]
		}

		if len(stack) < d {
			stack = append(stack, digit)
		}
	}

	return string(stack)
}
