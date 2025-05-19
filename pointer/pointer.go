package pointer

import (
	"fmt"
)

//--------------  Use case 1: pointer --------------------//
func PointerFirst() {
	x := 10

	var p *int = &x // p := &x   //✅ Go สามารถ infer type ให้อัตโนมัติ

	fmt.Println("Value of x:", x)  // Output: Value of x: 10
	fmt.Println("Value at p:", *p) // Output: Value at p: 10

	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20
}

func changValue(value *int) {
	*value = 10
}

func PointerMainCase1() {
	x := 20
	changValue(&x)
	fmt.Println(x)
}

//--------------  Use case 2: modified ข้อมูลต้นฉบับผ่าน Address --------------------//
type Employee struct {
	Name   string
	Salary int
}

func giveMoney(e *Employee, m int) {
	e.Name = "mark"
	e.Salary += m
}

func PointerMainCase2() {

	//------ การส่งอีกวิธีแบบเต็ม -----//
	//emp := Employee{Name: "John Doe", Salary: 50000}
	//fmt.Println("Befor raise:", emp)
	// giveMoney(&emp, 500)
	// fmt.Println("After raise:", emp)

	//----- การส่งอีกวิธีแบบสั้น -----//
	emp := &Employee{Name: "Alice", Salary: 2000}
	giveMoney(emp, 500)
	fmt.Println("After raise:", *emp)

}

//--------------  Use case 3: ดึงพวก config มาใช้ --------------------//
// Config represents the application configuration
type Config struct {
	LogLevel string
	Port     int
}

// UpdateConfig modifies the provided configuration
func UpdateConfig(c *Config, logLevel string, port int) {
	c.LogLevel = logLevel
	c.Port = port
}

func PointerMainCase3() {
	// Initial configuration
	appConfig := &Config{
		LogLevel: "info",
		Port:     8080,
	}

	fmt.Println("Initial Config:", appConfig)

	// Update configuration
	UpdateConfig(appConfig, "debug", 9000)
	fmt.Println("Updated Config:", appConfig)
}

//---------------------------------------------------//

//--------------  Use case 4: Pointer Receiver กับ Struct --------------------//

type Person struct {
	Name string
	Age  int
}

func (p *Person) Update(name string, age int) {
	p.Name = name
	p.Age = age
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func PointerMainCase4() {
	person := Person{Name: "Alice", Age: 25}

	// ใช้ method ที่แก้ไขค่าจริง
	person.Update("mark", 27)

	// ใช้ method ที่อ่านค่าปัจจุบัน
	person.Greet() // 👉 Hello, my name is Bob and I'm 30 years old.
}

//---------------------------------------------------//
func PointerMain() {
	// x := 20
	// changValue(&x)
	// fmt.Println(x)

	emp := Employee{Name: "John Doe", Salary: 50000}
	fmt.Println("Befor raise:", emp)

	giveMoney(&emp, 500)
	fmt.Println("After raise:", emp)

}
