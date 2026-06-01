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

func main() {
	topic, maxProblem, ok := parseArgs(os.Args)
	prog := filepath.Base(os.Args[0])
	if !ok {
		help.Print(prog)
		return
	}

	switch topic {
	case "i":
		problem.Print(maxProblem, interval.Ask)
	case "s":
		problem.Print(maxProblem, scale.Ask)
	default:
		help.Print(prog)
	}
}

func parseArgs(args []string) (string, int, bool) {
	ok := len(args) > 1
	topic := ""
	maxProblem := 0
	var err error
	for i, v := range args {
		switch i {
		case 1:
			topic = v
		case 2:
			maxProblem, err = strconv.Atoi(v)
			if err != nil {
				maxProblem = 0
			}
		}
	}
	return topic, maxProblem, ok
}
