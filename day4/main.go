package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ArrayMerge([]string{"alttera", "academy", "malang"}, []string{"academy", "jakarta"}) )
	fmt.Println(Mapping([]string{"alta", "malang", "alta", "jakarta", "malang"}))
	fmt.Println(munculSekali("1234321"))
} 

// Soal no 1
func ArrayMerge (arrayA []string, arrayB []string) []string {
	mergedArr := append(arrayA, arrayB...)
	uniqueArr := []string{}

	isFound := false
	for i, v := range mergedArr {
		if len(uniqueArr) == 0 {
			uniqueArr = append(uniqueArr, v)
		} else {
			isFound = false
			for j := i + 1; j < len(mergedArr); j++ {
				if (v == mergedArr[j]) {
					isFound = true
					// break
				}
			}
			
			if !isFound {
				uniqueArr = append(uniqueArr, v)
			}
		}
	}
	
	return uniqueArr
}



// Soal no 2
func Mapping(slice []string) map[string]int {
	var resultMapping = map[string]int{}

	for _, keyMap := range slice {
			resultMapping[keyMap]++
	}

	return resultMapping
}

// Soal No 3
func munculSekali(angka string) []int {
	tempMap := map[string]int{}
	result := []int{}
	for _, v := range angka {
		tempMap[string(v)]++
	}

	for key, v := range tempMap {
		if v == 1 {
			convertNumber, _ := strconv.Atoi(key)
			result = append(result, convertNumber)
		}
	}

	return result
}