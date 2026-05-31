package interval

import (
	"fmt"
	"math/rand"
)

func Ask(rd *rand.Rand) (string, string) {
	ok := false
	var lPitch, rPitch pitchPack
	var interval string
	for !ok {
		lPitch, rPitch = pitch(rd)
		interval, ok = answer(lPitch, rPitch)
	}

	question := fmt.Sprintf(
		"%c%c - %c%c",
		lPitch.pitch, lPitch.accidental(),
		rPitch.pitch, rPitch.accidental(),
	)
	return question, interval
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
