package main

import "fmt"

func main(){
  // membuat variabel a dan b
  var a bool = true
  var b bool = false

  // membuat instruksi logika
  if (a && b){
    fmt.Printf("kondisi a dan b adalah benar")
  }
  if (a || b){
    fmt.Printf("kondisi salah satu benar")
  }

  // mengganti value a dan b
  a = false
  b = true

  // membuat instruksi logika
  if (a && b){
    fmt.Printf("kondisi a dan b adalah benar")
  }else{
    fmt.Printf("tidak benar")
  }
  if(!(a && b)){
    fmt.Printf("kondisi benar")
  }
}

