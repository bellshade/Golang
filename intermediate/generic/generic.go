package main

import "fmt"

func main() {
	// Type Parameter
	resultString := Length[string]("Hello")
	fmt.Println(resultString)
	resultInt := Length[int](10)
	fmt.Println(resultInt)

	// Multiple Type Parameter
	MultipleTypeParameter[string, int]("Hello World", 100)

	// Type Comparable
	same := IsSame[string]("Hai", "Hai")
	fmt.Println(same)

	// Type Constraint
	manager := GetName[Manager](&MyManager{Name: "John"})
	fmt.Println(manager)

	// Type Sets
	fmt.Println(Min[int](1, 2))

	// Type Approximation
	fmt.Println(Min[Age](3, 4))

	// Generic Type
	names := Bag[string]{"Hello", "World"}
	PrintBag(names)

	// Generic struct
	DataStruct()

	// Generic Method
	GenericMethod()

	// Generic Interface
	Interface()

	// Inline Type Constraint
	fmt.Println(FindMin(1, 3))

	// Generic Type Parameter
	GetFirstTest()
}

// Untuk mengetahui apakah function tersebut adalah generic, kita perlu menambahkan Type Parameter pada function tersebut.
// Pembuatan Type Parameter menggunakan [] (kurung kotak) dan menuliskan nama type yang ingin kita gunakan.
// Biasanya nama untuk Type Parameter hanya menggunakan 1 huruf saja, contohnya: T, K, V, dan lain-lain.
func Length[T any](param T) T {
	return param
}

// Type Parameter bisa digunakan lebih dari 1, apabila kita ingin menggunakan lebih dari 1 type parameter.
// Maka cukup gunakan tanda koma (,) untuk menandakan sebagai pemisah antar type parameter.
// Nama Type Parameter tidak boleh sama dengan Type Parameter yang telah didefinisikan sebelumnya.
func MultipleTypeParameter[T1, T2 any](param T1, param2 T2) {
	fmt.Println(param)
	fmt.Println(param2)
}

// Comporable merupakan interface atau any yang diimplementasikan oleh type data yang bisa dibandingkan(menggunakan operator == dan !=).
// Contohnya: Booleans, Numbers, Strings, Pointers, Channels, Interfaces, arrar yang isinya ada comparable type, atau struct yang isinya ada comparable type.
func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

// Jika kita membuat sebuah type yang sesuai dengan kontrak interface, maka dianggap sebagai implementasi interface tersebut.
// Type Parameter juga dapat digunakan untuk mengimplementasikan interface.
// Kita bisa gunakan constraint dengan menggunakan interface.
// Secara otomatis, semua interface yang compatible dengan type constraint tersebut bisa digunakan
type Employee interface {
	GetName() string
}

func GetName[T Employee](employee T) string {
	return employee.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// Type sets merupakan sebuah interface dan salah satu fitur dari golang generic
// Fitur ini membuat kita bisa menentukan lebih dari satu type contraint yang diperbolehkan pada type parameter
// Type sets hanya bisa digunakan pada type parameter, tidak bisa dipakai sebagai type data field ataupun variable

// Type Approximation merupakan fitur golang yang dimana kita bisa menyebutkan bahwa semua contraint dengan tipe tersebut
// dan juga yang memiliki dasarnya adalah tipe tersebut, maka bisa digunakan.
// Untuk menggunakan Type Approximation kita bisa menggunakan tanda ~ (tilde)
type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

// Generic type
type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

// struct juga mendukung generic
// Dengan menggunakan generic, kita bisa membuat field dengan tipe data yang sesuai dengan type paramater
type Data[T any] struct {
	First  T
	Second T
}

func DataStruct() {
	data := Data[string]{
		First:  "Hai",
		Second: "Dunia",
	}
	fmt.Println(data)
}

// selain difunction, generic juga bisa digunakan di method, namun generic dimethod merupakan generic yang terdapat di struct nya.
// Kita wajib menyebutkan semua type parameter yang digunakan pada struct tersebut. apabila tidak digunakan kita bisa menggunakan _ (underscore)
// sebagai tanda bahwa type parameter tersebut tidak digunakan.
// Method tidak bisa memiliki type parameter yang mirip dengan difunction.
func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func GenericMethod() {
	data := Data[string]{
		First:  "Hai",
		Second: "Dunia",
	}

	fmt.Println(data.SayHello("World"))
	fmt.Println(data.ChangeFirst("Hello"))
}

// Generic bisa digunakan di interface
// Secara otomatis, semua struct yang ingin mengikuti interface tersebut harus menggunakan generic juga
type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

type MyData[T any] struct {
	Value T
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func Interface() {
	myData := MyData[string]{}
	result := ChangeValue[string](&myData, "Helo")
	fmt.Println(result)
}

// Inline Type Contraint
// Tidak perlu menggunakan type sets ataupun type declaration, kita bisa menggunakan inline type constraint
func FindMin[T interface{ int | int16 | int32 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

// Generic di Type Parameter
func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func GetFirstTest() {
	names := []string{
		"Hai", "Dunia",
	}

	first := GetFirst[[]string, string](names)
	fmt.Println(first)
}
