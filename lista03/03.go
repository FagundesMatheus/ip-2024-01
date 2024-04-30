package main
import "fmt"

func main() {
  var n int
  fmt.Scanln(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
  }
  for i := n - 1; i >= 0; i--{
      fmt.Print(nums[i])
      fmt.Print(" ")
  }
}
