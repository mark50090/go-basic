package iteration

import (
	"fmt"
)

func Iterationforloop() {
	for i := 1; i <= 10; i++ {
		fmt.Println("number is ", i)
	}
}

func Iterationdowhileloops() { //ทำก่อนค่อยเช็คเงื่อนไข
	var number int = 10
	for {
		fmt.Println("number is ", number)
		number++
		if number <= 10 {
			break
		}

	}
}

func Iterationwhileloops() { //ทำก่อนค่อยเช็คเงื่อนไข
	var number int = 10
	i := 1
	for i <= number {
		fmt.Println("number is ", i)
		i++
	}
}
