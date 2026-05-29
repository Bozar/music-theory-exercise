package interval

import (
	"fmt"
	"math/rand"
	"time"
)

func Print(countProblem int) {
	var ask, answer string
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	wrong := []string{}
	if countProblem < minProblem {
		countProblem = maxProblem
	}
	countProblem = min(countProblem, maxProblem)

	for i := 1; i < countProblem+1; i++ {
		ask, answer = problem(i, rd)
		if verify(answer) {
			continue
		}
		wrong = append(wrong, fmt.Sprintf("%s | %s", ask, answer))
	}

	fmt.Println()
	for _, v := range wrong {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("$$: %d/%d\n", countProblem-len(wrong), countProblem)
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

func problem(count int, rd *rand.Rand) (string, string) {
	ok := false
	var lPitch, rPitch pitchPack
	var interval string
	for !ok {
		lPitch, rPitch = pitch(rd)
		interval, ok = answer(lPitch, rPitch)
	}

	ask := fmt.Sprintf(
		"%02d: %c%c - %c%c", count,
		lPitch.pitch, lPitch.accidental(),
		rPitch.pitch, rPitch.accidental(),
	)
	fmt.Printf("%s\n", ask)
	return ask, interval
}

func pitch(rd *rand.Rand) (pitchPack, pitchPack) {
	lIndex, rIndex := index(rd)
	lQuality, rQuality := quality(rd)

	lPitch := pitchPacks[lIndex]
	rPitch := pitchPacks[rIndex]
	lPitch.quality = lQuality
	rPitch.quality = rQuality
	return lPitch, rPitch
}

func index(rd *rand.Rand) (int, int) {
	left := rd.Intn(maxPitchPacksIndex)
	right := left
	for right == left {
		right = rd.Intn(maxPitchPacksIndex)
	}
	if left > right {
		right += maxPitchPacksIndex
	}
	//	fmt.Printf("%d,%d\n", left, right)
	return left, right
}

func quality(rd *rand.Rand) (int, int) {
	left := rd.Intn(maxStep-minStep) + minStep
	right := rd.Intn(maxStep-minStep) + minStep
	return left, right
}

func answer(lPitch pitchPack, rPitch pitchPack) (string, bool) {
	size := rPitch.size - lPitch.size + 1
	step := (rPitch.halfStep + rPitch.quality) -
		(lPitch.halfStep + lPitch.quality)
	quality := stepToQuality[size*100+step]
	return fmt.Sprintf("%c%d", quality, size), (quality != 0)
}
