package main
import "fmt"

func main() {
   var str string;
   var num int;
   for{
       fmt.Scanln(&num)
       recursiva(num, str)
       fmt.Println(str)
       str = ""
   }
}
func recursiva(num int, str string){
    
    if num == 1{str += "1"; return}
    if num == 0{str += "0"; return}
    str += string(num%2)
    str2 := str
    recursiva(num/2, str2)
}
