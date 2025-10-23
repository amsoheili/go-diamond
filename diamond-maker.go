package main

import "fmt"

func main() {
	maxWidth := 101
	fmt.Println(" ") // 50 49 48
	// stars 1 - 3 - 5 - .... - 101 - 99 - 97 - ....  - 1
	var isUpperTriangle bool = true

	spaceNo := (maxWidth - 1) / 2

	starNo := 1

	for i := 0; i < maxWidth; i++ {
		for spi := 0; spi < spaceNo; spi++ {
			fmt.Print(" ")
		}
		for sti := 0; sti < starNo; sti++ {
			fmt.Print("*")
		}
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
		fmt.Println()
	}
}
