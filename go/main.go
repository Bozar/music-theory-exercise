package main

import (
	"fmt"
)

func main() {
	var input, check string
	count := 0
	maxCount := 3
	for i := 1; i < maxCount+1; i++ {
		check = fmt.Sprintf("m%d", i)
		fmt.Printf("answer: %s\n$?: ", check)
		fmt.Scan(&input)
		if input != check {
			fmt.Println("wrong")
		} else {
			count += 1
		}
	}
	fmt.Printf("$$: %d/%d\n", count, maxCount)
}
