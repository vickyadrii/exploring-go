package main

import "fmt"

func main() {
	fmt.Println(ArrayMerge([]string{"alttera", "academy", "malang"}, []string{"academy", "jakarta"}) )
}

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