
package main
import "fmt"

func main() {MaxNum()}
func MaxNum(){
  var n, max, index int
  max = 0
  index = 0
  fmt.Scanln(&n)
  if n != 0{
  nums := make([]int, n)
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
      if nums[i] > max {max = nums[i]; index = i;}
  }
  fmt.Printf("%d %d \n", index, max)
  MaxNum()

}
}
