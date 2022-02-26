package array

import "fmt"

// In Go Array is value. It's more recommended to use slice in Go,
// because it's style is idomatic Go.

func Array() {
	// BASIC ARRAY DEFINITION AND ASSIGNMENT
	var anArray [5]int32
	anArray[0] = 0x48 // H
	anArray[1] = 0x45 // E
	anArray[2] = 0x4c // L
	anArray[3] = 0x4c // L
	anArray[4] = 0x4f // O

	var secArray [5]byte = [5]byte{87, 79, 82, 76, 68}

	fmt.Printf("%c", anArray)
	fmt.Printf("%c\n", secArray)

	// SUM NUMBERS
	myNumbers := [2]int{9, 184}
	fmt.Println("myNumbers       :", myNumbers)
	fmt.Println("myNumbers[0]    :", &myNumbers[0])
	fmt.Println("myNumbers[1]    :", &myNumbers[1])
	fmt.Println("sum of myNumbers:", Sum(&myNumbers))

	// ANOTHER SUM NUMBERS
	myNumbersSec := new([2]int)
	myNumbersSec[0] = 90
	myNumbersSec[1] = 10

	fmt.Println("myNumbers       :", myNumbersSec)
	fmt.Println("myNumbers[0]    :", &myNumbersSec[0])
	fmt.Println("myNumbers[1]    :", &myNumbersSec[1])
	fmt.Println("sum of myNumbers:", Sum(myNumbersSec))

}

// Use pointer for C like-behaviour and efficiency
// note: this style is not recommended
func Sum(numbers *[2]int) (total int) {
	for _, num := range *numbers {
		total += num
	}

	return
}
