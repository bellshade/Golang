# maps

Golang menyediakan tipe data lainnya bernama **map** yang memetakan key unik ke nilai. key adalah objek yang kita gunakan untuk mengambil nulai di kemudian nanti.diberikan key dan nilai, kita dapat menyimpan nilai dalam objek map. setelah nilai disimpan, kita dapat mengambilnya dengan menggunakan key.

## mendifinisikan map

kita harus menggunakna fungsi **make** untuk membuat peta

```golang
/* deklarasi variabel dengan value nil */
var nama_variabel map[tipe_data_key]value_dari_tipe_data

/* definisi map sebagai nil map yang tidak assign value apapun */
var nama_variabel = make(map[tipe_data_key]value_dari_tipe_data)
```
## menghapus key pada map

``delete()`` digunakan untuk menghapus entri dari map. ini membutuhkan map dan key yang sesuai yang akan dihapus

contoh sederhana dari penggunaan map

```golang
package main

import "fmt"

func main(){
  var namaIbukota map[string]string
  namaIbukota = make(map[string]string)

  namaIbukota["indonesia"] = "jakarta"
  namaIbukota["jepang"] = "tokyo"
  namaIbukota["india"] = "nwe delhi"

  // print map menggunakan key
  for negara := range namaIbukota{
    fmt.Println("ibukota ", negara, :"adalah", namaIbukota[country])
  }

  // mengecek apakah adakah negara ibukota yang termasuk
  ibuKota, ok := namaIbukota["jerman"]

  if (ok){
    fmt.Println("ibuka dari jerman adalah", ibuKota)
  }
  else{
    fmt.Println("ibukota dari negara jerman tidak termasuk")
  }  
}
```
