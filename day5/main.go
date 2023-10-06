package main

import "fmt"


// Soal 1
type Car struct {
	Type string
	FullIn float32	
}

func (car Car) predictDistance() float32 {
	distance := car.FullIn / 1.5
	return distance
}

func main() {
	// Soal 1
	var car Car
	car.FullIn = 2
	distance := car.predictDistance()
	fmt.Println(distance)
}