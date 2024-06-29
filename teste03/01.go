package main

import "fmt"

func main()  {
	for{	
		var str string
		fmt.Scanln(&str)
		if str == "#"{
			return
		}
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	}