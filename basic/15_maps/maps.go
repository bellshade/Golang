package main

import "fmt"

func main(){
  var ibuKotaNegara map[string]string
  // membuat map
  ibuKotaNegara = make(map[string]string)

  ibuKotaNegara["indonesia"] = "jakarta"
  ibuKotaNegara["jepang"] = "tokyo"
  ibuKotaNegara["perancis"] = "paris"
  ibuKotaNegara["italia"] = "roma"

  // print dengan menggunakan key
  for negara := range ibuKotaNegara{
    fmt.Println("ibukota dari", negara, "adalah", ibuKotaNegara[negara])
  }

  // cek apakah suatu ibukota termasuk ke dalam map
  ibuKota, ok := ibuKotaNegara["jerman"]

  // jika terdapat maka akan menampilkan
  if (ok){
    fmt.Println("ibukota dari negara jerman adalah", ibuKota)
  } else {
    fmt.Println("ibukota dari negara jerman tidak termasuk")
  }
}

