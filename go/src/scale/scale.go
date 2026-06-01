package scale

import (
	"fmt"
	"math/rand"
)

func Ask(rd *rand.Rand) (string, string) {
	sz := size(rd)
	lIndex, rIndex := index(sz, rd)

	question := fmt.Sprintf("%c %+d", pitches[lIndex], sz)
	answer := pitches[rIndex]
	return question, string(answer)
}

func index(sz int, rd *rand.Rand) (int, int) {
	left := rd.Intn(maxSize)
	right := left
	lenP := len(pitches)

	if sz > 0 {
		right += sz - 1
	} else {
		right += sz + 1
	}
	if right >= lenP {
		right -= lenP
	} else if right < 0 {
		right += lenP
	}
	return left, right
}

func size(rd *rand.Rand) int {
	sz := rd.Intn(maxSize-minSize) + minSize
	if rd.Intn(2) == 0 {
		sz = -sz
	}
	return sz
}
