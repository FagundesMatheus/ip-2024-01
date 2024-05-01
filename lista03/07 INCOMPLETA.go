
package main
import "fmt"

func main() {
  var n, max int
  fmt.Scanln(&n)
  nums := make([]int, n)
  var m map[int]int
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
      if nums[i] > max {max = nums[1]}
      m[nums[i]] += m[nums[1]]
  }
  for i := 0; i < max; i++{
      fmt.Println("(%d) %d", i, m[i])
  }
 
}
