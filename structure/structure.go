package structure

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	Name   string
	Weight int
	Height int
	Grade  string
}

func StructureArray() {
	var a [3]int = [3]int{1, 2, 3}
	var b [2]string = [2]string{"Go", "Lang"}
	var c [2]Person = [2]Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}
	mark := [3]int{7, 8, 9}

	// var nums [3]int         // array ของ int
	// var names [2]string     // array ของ string
	// var flags [4]bool       // array ของ bool
	// var floats [2]float64   // array ของ float64

	fmt.Println(mark)
	fmt.Println(b)
	fmt.Println(c)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}

func StructureSlice() {
	//var names []string
	var mark []int = []int{4, 5, 6}
	kk := []int{14, 15, 16}

	// names = append(names, "mark")
	fmt.Println(mark, kk)
}

func StructureMap() {
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])

	// Update the value for a key
	myMap["banana"] = 12

	// Delete a key-value pair
	delete(myMap, "orange")

	// Iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	// Checking if a key exists
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found inmap")
	}
}

func StructureStruc() {
	var mm []Student = []Student{
		{Name: "KRITTAPAS",
			Weight: 65,
			Height: 171,
			Grade:  "2.99"},
	}

	var studentData Student
	studentData.Name = "mark"
	mm[0].Name = "mark01"

	nark := Person{
		Name: "mark",
		Age:  25,
	}

	u, _ := json.Marshal(Person{Name: "Bob", Age: 10})

	//fmt.Println(string(u))
	fmt.Println(mm, studentData, nark, string(u))
}
