package main
import(
  "fmt"
  "strings"
)

// make a alias
type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string{
  var result []string
  for _, each := range data{
    if filtered := callback(each); filtered{
      result = append(result, each)
    }
  }
  return result
}

func main(){
  var data = []string{
    "wick",
    "jason",
    "ethan",
  }
  var dataContains0 = filter(data, func(each string) bool{
    return strings.Contains(each, "o")
  })
  var dataLength5 = filter(data, func(each string) bool{
    return len(each) == 5
  })

  fmt.Println("data asli \t\t:",data)
  // data asli : [wick jason ethan]

  fmt.Println("Filter ada huruf \"o\"\t:", dataContains0)
  // filter ada huruf "o" : [jason]

  fmt.Println("Filter jumlah huruf \"S\"\t:", dataLength5)
  // filted jumlah huruf "5" : [jason ethan]

}

// string.contains =  (string dicek, string cek) return will be boolean
// var result = strings.Contains("Golang", "ang")
