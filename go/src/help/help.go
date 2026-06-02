package help

import (
	"fmt"
)

func Print(prog string) {
	fmt.Printf("%s {i|s|v} [max_problem=10]\n\n", prog)
	fmt.Println("i: Interval exercise")
	fmt.Println("s: Scale exercise")
	fmt.Println("v: Show version")
}

func Version() {
	fmt.Println(version)
}
