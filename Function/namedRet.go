package main
import "fmt"

func split(sum int) (x,y int) {
  x = sum * 4 / 9  //7.5
  y = sum - x  //10
  return
}

func main(){
  fmt.Println(split(17))
}
