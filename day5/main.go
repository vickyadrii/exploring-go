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


// Soal 2
type Student struct {
	name []string
	score []int
}


func main() {
	// Soal 1
	var car Car
	car.FullIn = 2
	distance := car.predictDistance()
	fmt.Println(distance)


	// Soal 2
	var students Student
	for i := 0; i <= 4; i++ {
		students.name = append(students.name, "")
		students.score = append(students.score, 0)

		fmt.Println("Masukkan nama murid ke ", i+1, " : ")
		fmt.Scanln(&students.name[i])
		
		fmt.Println("Masukkan nilai murid ke ", i+1, " : ")
		fmt.Scanln(&students.score[i])
	}

	fmt.Println(students)
	
}