package help

import (
	"fmt"
	"os"
	"path/filepath"
)

func Print() {
	prog := filepath.Base(os.Args[0])
	fmt.Printf("%s {i|s} [max_problem]\n\n", prog)
	fmt.Println("i: Interval")
	fmt.Println("s: Scale")
}
