package main
import "fmt"

func calculate(params ...int)(int){
  var total int = 0
  for _,number:=range params{
    total+=number
  }
  return total
}

func main(){
  fmt.Println(calculate(3,2,4,5,2,4,5))
}
