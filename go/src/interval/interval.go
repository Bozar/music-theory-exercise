package interval

import (
	"fmt"
	"math/rand"
	"time"
)

func Print() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < maxProblem+1; i++ {
		problem(i, rd)
	}
}

func problem(count int, rd *rand.Rand) {
	lPitch, rPitch := pitch(rd)

	fmt.Printf(
		"%02d: %c%c - %c%c\n", count,
		lPitch.pitch, lPitch.accidental(),
		rPitch.pitch, rPitch.accidental(),
	)
}

func pitch(rd *rand.Rand) (pitchPack, pitchPack) {
	lIndex, rIndex := index(rd)
	lStep, rStep := step(rd)

	lPitch := pitchPacks[lIndex]
	rPitch := pitchPacks[rIndex]
	lPitch.halfStep = lStep
	rPitch.halfStep = rStep
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

func step(rd *rand.Rand) (int, int) {
	left := rd.Intn(maxStep-minStep) + minStep
	right := rd.Intn(maxStep-minStep) + minStep
	return left, right
}

//func Print() {
//	fmt.Printf("%c-%d\n", pitchPacks[0].pitch, pitchPacks[0].halfStep)
//	var input, check string
//	count := 0
//	maxCount := 3
//	for i := 1; i < maxCount+1; i++ {
//		check = fmt.Sprintf("m%d", i)
//		fmt.Printf("answer: %s\n$?: ", check)
//		fmt.Scanln(&input)
//		//		fmt.Scanf("%s\n", &input)
//		//		fmt.Scan(&input)
//		if input != check {
//			fmt.Println("wrong")
//		} else {
//			count += 1
//		}
//	}
//	fmt.Printf("$$: %d/%d\n", count, maxCount)
//}
