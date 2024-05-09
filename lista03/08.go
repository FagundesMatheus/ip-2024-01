package main

import "fmt"

func main() {
    var num int
    for fmt.Scanln(&num) != 0 {
        binary := recursive(num)
        fmt.Println(binary)
    }
}

func recursive(num int) string {
    if num == 0 {
        return "0"
    }
    if num == 1 {
        return "1"
    }
    return recursive(num/2) + string(num%2)
}
