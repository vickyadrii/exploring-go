package main

import "fmt"

func add(a, b int)int {

	result := a+b
	fmt.Println(result)
	return result

}

func main() {
	add(1, 2)

	// Array
		// Banyak data
		// 1 tipe
		// index
		// ukuran fix

	scoreData := [5]int{1,2,3,4,5}

	for i := 0; i < len(scoreData); i++ {
		fmt.Println(scoreData[i]);
	}

	fmt.Println(scoreData)

	// Slice
		// 	Banyak data
		// 	index
		// 	dinamis
	

	newScore := make([]int, 10)
	var score []int = []int{1,2,3}

	fmt.Println(newScore)
	score = append(score, 10) // Tambah data menggunakan append
	score[0] = 104
	fmt.Println(score)


	// Map
		// banyak data
		// key value
		// dinamis
		// key unique
	

	var dataPenjualan = make(map[string]int)

	dataPenjualan["Januari"] = 10000
	dataPenjualan["Februari"] = 20000
	dataPenjualan["Februari"] = 10000
	dataPenjualan["Maret"] = 42000

	_, isFound := dataPenjualan["Februari"] // untuk mencari key apakah ada atau tidak ada
	fmt.Println(dataPenjualan)
	fmt.Println(dataPenjualan["Januari"])
	fmt.Println(isFound)

}
