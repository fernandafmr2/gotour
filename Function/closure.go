package main
import "fmt"

func main(){
  var getminmax = func(n []int)(int,int){
    var min,max int
    for i, e:=range n{
      switch{
      case i == 0:
        max,min =e,e
      case e>max:
        max = e
      case e<min:
        min = e
      }
    }
    return min,max
  }
  var numbers = []int{2,4,2,5,76,4}
  var min,max = getminmax(numbers)
  fmt.Printf("Data : %v\nmin : %v\nmax : %v\n",numbers, min, max)
}
