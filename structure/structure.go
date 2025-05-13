package structure

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
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
