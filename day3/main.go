package main

import "fmt"

func add(a, b int)int {

	result := a+b
	fmt.Println(result)
	return result

}

func main() {
	add(1, 2)
}
