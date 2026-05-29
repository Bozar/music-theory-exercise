package main

import (
	"github.com/Bozar/music-theory-exercise/go/src/help"
	"github.com/Bozar/music-theory-exercise/go/src/interval"

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
