package main

import "fmt"

func main() {
	// variables

	// Strings
	var firstName string = "Vicky"
	var middleName = "Herdiansyah"
	var lastName string

	fmt.Println(firstName, middleName, lastName)

	firstName = "Jindan"
	lastName = "Za"

	fmt.Println(firstName, middleName, lastName)

	nameOne := "Nopal";
	fmt.Println(nameOne)

	// Ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits & Memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 120 // Stores unsigned integers, (Tidak menerima bilangan minus atau < 0)

	fmt.Println(numOne, numTwo, numThree)

	// Floats
	var scoreOne float32 = 10.43
	var scoreTwo float64 = 746745645.43543
	scoreThree := 234234.432

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
