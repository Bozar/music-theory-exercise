package main

import (
	"github.com/Bozar/music-theory-exercise/go/src/help"
	"github.com/Bozar/music-theory-exercise/go/src/interval"
	"github.com/Bozar/music-theory-exercise/go/src/problem"
	"github.com/Bozar/music-theory-exercise/go/src/scale"

	"os"
	"path/filepath"
	"strconv"
)

type args struct {
	topic      string
	maxProblem int

	prog string
	ok   bool
}

func main() {
	pa := parseArgs(os.Args)
	if !pa.ok {
		help.Print(pa.prog)
		return
	}

	switch pa.topic {
	case "i":
		problem.Print(pa.maxProblem, interval.Ask)
	case "s":
		problem.Print(pa.maxProblem, scale.Ask)
	case "v":
		help.Version()
	default:
		help.Print(pa.prog)
	}
}

func parseArgs(input []string) args {
	pa := args{}
	pa.ok = len(input) > 1
	var err error
	for i, v := range input {
		switch i {
		case 0:
			pa.prog = filepath.Base(v)
		case 1:
			pa.topic = v
		case 2:
			pa.maxProblem, err = strconv.Atoi(v)
			if err != nil {
				pa.maxProblem = 0
			}
		}
	}
	return pa
}
