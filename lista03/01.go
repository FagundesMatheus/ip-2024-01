package main
import "fmt"

func main() {
  var n, num int
  fmt.Scanln(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
  }
  for i := 0; i < n; i++{
      fmt.Scan(&num)
      for i := 0; i < n; i++{
          if nums[i] == num {
              fmt.Println("ACHEI") 
              break
          }
          if nums[i] != num && i == n - 1{fmt.Println("NÃO ACHEI")}
      }
  }
  
}
