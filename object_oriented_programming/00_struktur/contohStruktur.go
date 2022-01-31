package main

import "fmt"

// deklarasi struct
type Motor struct {
  pemilik string
  merk string
  nomorMesin int
}

// fungsi untuk print
func (motor Motor) tampilkan_detail() {
  fmt.Printf("pemilik motor %s adalah %s\n", motor.merk, motor.pemilik)
  fmt.Printf("dengan nomor mesin %d\n", motor.nomorMesin)
}

func main() {
  // deklarasi struct
  motor1 := Motor{"james", "yamaha", 27413}
  
  // printing detail dari motor
  motor1.tampilkan_detail()

  // modifiikasi value
  motor1.pemilik = "majed"
  motor1.nomorMesin = 27582

  // print detaul dari yang sudah diryubah
  motor1.tampilkan_detail()
}

