package main

import "fmt"

// membuat structure person1 dengan field name,age,speed
type Person1 struct {
	name  string
	age   int
	speed int
}

// membuat structure person2 dengan field yang sama tetapi ada tambahan property skin
type Person2 struct {
	name  string
	age   int
	speed int
	skin  string
}

// membuat interface human yang siapa mau mengimplementasikannya harus mempunyai method walk
type Human interface {
	walk()
}

// membuat method getSkin khusus untuk struct person2
func (p2 *Person2) getSkin() string {
	return p2.skin
}

// meingimplementasikan interface human untuk struct person2
func (p2 *Person2) walk() {
	p2.speed++
	fmt.Printf("speed after walk: %d\n",p2.speed)
}

// mengimplementasikan interface human untuk struct person1
func (p1 *Person1) walk() {
	p1.speed++
	fmt.Printf("speed after walk: %d\n", p1.speed)
}

// membuat instance human yang memiliki parameter name, age, skin lalu mereturn interface Human
func NewHuman(name string, age int, skin string) Human {
	// pengecekan jika nilai argumen skin itu ada isinya atau tidak sama dengan "" maka kita menggunakan person2

	if skin != "" {
		return &Person2{name: name, age: age, skin: skin}
	} else {
		// sedangkan jika parameter skin itu "" maka kita menggunakan structure person1
		return &Person1{name: name, age: age}
	}
}

func main() {
	/*
		 	membuat instance human dengan nama = "your name", age = 20 dan skin = "white"
			jika kalian tidak ingin memasukkan nilai skin kalian bisa beri aja nilai ""
	*/
	human := NewHuman("your name", 20, "white")
	
	/*
		mengkonversi dan melakukan pengecekan apakah konversi tipe interface Human ke tipe *Person1 gagal atau berhasil jika ok maka eksekusi kode if ini
	*/
	if person1, ok := human.(*Person1); ok {
		// print value dari person1
		fmt.Printf("name: %s\n age: %d\n speed: %d\n", person1.name, person1.age, person1.speed)
		// tampilkan speed person1 setelah walk saat ini
		person1.walk()
	}
	/*
		mengkonversi dan melakukan pengecekan apakah konversi tipe interface Human ke tipe *Person2 gagal atau berhasil jika ok maka eksekusi kode if ini
	*/
	if person2, ok := human.(*Person2); ok {
		// print value dari person2, disini kita bisa memakai method getSkin() dari person2, karena kita sudah mengkonversinya, jadi kita bisa akses metode tersebut
		fmt.Printf("name: %s\n age: %d\n speed: %d\n skin: %s\n", person2.name, person2.age, person2.speed, person2.getSkin())
		// tampilkan speed person2 setelah walk saat ini
		person2.walk()
	}
}
