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
	ok := false
	var lPitch, rPitch pitchPack
	var interval string
	for !ok {
		lPitch, rPitch = pitch(rd)
		interval, ok = answer(lPitch, rPitch)
	}

	fmt.Printf(
		"%02d: %c%c - %c%c\n", count,
		lPitch.pitch, lPitch.accidental(),
		rPitch.pitch, rPitch.accidental(),
	)
	fmt.Printf("%s\n", interval)
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
