package main

import "fmt"

func main() {
	maxWidth := 101

	var isUpperTriangle bool = true

	spaceNo := (maxWidth - 1) / 2

	starNo := 1

	arrayMap := make([][]int, maxWidth)

	for i := 0; i < maxWidth; i++ {
		arrayMap[i] = []int{spaceNo, starNo}
		if starNo == maxWidth {
			isUpperTriangle = false
		}
		if isUpperTriangle {
			spaceNo--
			starNo += 2
		} else {
			spaceNo++
			starNo -= 2
		}
	}

	for i := 0; i < len(arrayMap); i++ {
		for j := 0; j < arrayMap[i][0]; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < arrayMap[i][1]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
