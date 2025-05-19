package functional

import (
	"fmt"
)

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Persons struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Persons) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func ParameterInterface() {
	dog := Dog{Name: "Buddy"}
	person := Persons{Name: "Alice"}

	makeSound(dog)
	makeSound(person)
}
