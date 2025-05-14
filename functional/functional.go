package functional

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	Firstname string
	Lastname  string
}

// การสร้าง method โดย (s Student) จะกลายเป็น method ทันที
func (s Student) FullName() string { //ให้ s เป็นตัวแทนของ struct เพื่อไปสร้าง intance ออกมาใหม่
	return s.Firstname + " " + s.Lastname
}

func ParameterFunction(name string, age int) string {
	data := Person{
		Name: "Bob",
		Age:  10,
	}
	// Marshal เป็น JSON
	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes)
}

func ParameterFunction2() {
	result := ParameterFunction("mark", 27)
	fmt.Println("parameterFunction", result)
}

func ParameterCallMethod() {
	student := Student{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}

	fullName := student.FullName()
	fmt.Println("Full Name", fullName)
}
