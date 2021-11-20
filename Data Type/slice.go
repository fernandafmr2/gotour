package main
import "fmt"

func main(){
  var str = []string{"Akame","Kitsune","Rem","Ram","Sagiri"}

  // slcing reference // MENGUBAH ISI ARRAY YANG DI REFERENCE
  // array[low:high] -> normal tapi di akhir-1
  // array[low:]    -> dari low sampai akhir
  // array[:high]    -> start index 0 sampai akhir-1
  // array[:]       -> dari index 0 sampai akhir
  var sliceStr = str[0:3]
  
  sliceStr[0] = "Remu"

  fmt.Println(str)
  fmt.Println(sliceStr)
  fmt.Println(len(sliceStr))
  fmt.Println(cap(sliceStr))
}


// function for slice 
// len(slice) -> get length slice
// cap(slice) -> get capacity slice 
// append(slice,data) -> make a new slice if slice before fuel
// make([]typeData, length, capacity)  -> Make a new slice 
// copy(destination,source) -> copy slice
