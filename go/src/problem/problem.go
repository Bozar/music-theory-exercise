package problem

import (
	"fmt"
	"math/rand"
	"time"
)

type Ask func(rd *rand.Rand) (string, string)

func Print(countProblem int, ask Ask) {
	countProblem = fixCount(countProblem, minProblem, maxProblem)
	var question, answer string
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	wrong := []string{}

	for i := 1; i < countProblem+1; i++ {
		question, answer = ask(rd)
		question = fmt.Sprintf("%02d: %s", i, question)
		fmt.Printf("%s\n", question)
		if verify(answer) {
			continue
		}
		wrong = append(wrong, fmt.Sprintf("%s | %s", question, answer))
	}

	fmt.Println()
	for _, v := range wrong {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("$$: %d/%d\n", countProblem-len(wrong), countProblem)
}

func fixCount(count int, minProblem int, maxProblem int) int {
	if count < minProblem {
		count = maxProblem
	}
	count = min(count, maxProblem)
	return count
}

func verify(answer string) bool {
	var input string
	fmt.Printf("$?: ")
	fmt.Scanln(&input)

	if input == answer {
		return true
	}
	fmt.Println("Wrong")
	return false
}
