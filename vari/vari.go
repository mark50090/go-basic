package vari

import (
	"fmt"
)

func Varitype() {
	// var age int = 30
	// var name string = "John"
	// var price float64 = 99.99
	// var isActive bool = true

	age := 30        // int
	name := "John"   // string
	price := 99.99   // float64
	isActive := true // bool

	fmt.Println(age, name, price, isActive)
}

func Constant() {
	const Pi = 3.14
	const AppName string = "MyApp"
	fmt.Println(Pi, AppName)
}

func Pointer() {
	var ptr *int
	var number = 10
	ptr = &number

	fmt.Println(ptr, number)
}

func Array() {

}

func Mainvari() {
	Varitype()
	Constant()
	Pointer()
}
