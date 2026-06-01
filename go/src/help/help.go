package help

import (
	"fmt"
)

func Print(prog string) {
	fmt.Printf("%s {i|s} [max_problem=10]\n\n", prog)
	fmt.Println("i: Interval")
	fmt.Println("s: Scale")
}
