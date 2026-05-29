package interval

type pitchPack struct {
	pitch    rune
	size     int
	quality  int
	halfStep int
}

func (pp *pitchPack) accidental() rune {
	switch pp.quality {
	case -1:
		return 'b'
	case 1:
		return '#'
	default:
		return 0
	}
}
