package condition

import (
	"fmt"
)

func Conditionifelse() {
	const number int = 56
	var grade string
	if number >= 80 {
		grade = "A"
		fmt.Println(grade)
	} else if number >= 70 {
		grade = "B"
		fmt.Println(grade)
	} else {
		grade = "C"
		fmt.Println(grade)
	}

	num1 := 5
	num2 := 10

	if sumNum := num1 + num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}

func Conditionswitch() {
	var dayOfWeek = 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

}

func Conditionswitchcondition() {
	var score int = 62
	var grade string

	switch {
	case score >= 80:
		grade = "A"
	case score >= 70:
		grade = "B"
	case score >= 60:
		grade = "C"
	default:
		grade = "F"
	}

	fmt.Printf("Your grade is %s", grade)
}
