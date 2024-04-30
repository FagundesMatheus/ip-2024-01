
package main
import "fmt"

func main() {
  
  funcao()
}
func funcao(){
  var k, n, count int
  fmt.Scanln(&n)
  if n >= 1 && n <= 1000{
      nums := make([]int, n)
      
      for i := 0; i < n; i++{
          fmt.Scan(&nums[i])}
          fmt.Scanln(&k)
      for i := 0; i < n; i++{
          if nums[i] >= k {count++}
          }
  }else{funcao()}
  fmt.Print(count)
}
