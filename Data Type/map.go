package main
import "fmt"

func main(){
  clos := map[string]string{
    "Name" : "Akame",
    "Anime" : "Akame ga kill",
  }

  delete(clos,"Anime")
  fmt.Println(clos)
  fmt.Println(clos["Name"])
  fmt.Println(clos["Anime"])
}


// function map 
// len(map)  -> Mendapatkan jumlah data map 
// map(key)  -> Mengambil data di map dengan key 
// map[key]=value  -> mengubah data di map dengan key 
// make(map[typeKey]typeValue)  -> membuat map baru
// delete(map,key) -> menghapus data di map dengan key
