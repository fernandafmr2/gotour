package main 
import "fmt"

func swap(arg1, arg2 string) (string,string){
	return arg2, arg1
}

func main(){
	a, b := swap("Pocari","Sweat")
	fmt.Println(a,b)
}