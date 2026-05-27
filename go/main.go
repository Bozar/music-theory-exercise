package main

import (
	"github.com/Bozar/mygo/src/help"
	"github.com/Bozar/mygo/src/interval"

	"os"
)

func main() {
	if len(os.Args) < 2 {
		help.Print()
		return
	}

	switch os.Args[1] {
	case "i":
		interval.Print()
	default:
		help.Print()
	}
}
