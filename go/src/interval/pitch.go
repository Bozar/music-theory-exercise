package interval

type pitchPack struct {
	pitch    rune
	interval int
	halfStep int
}

func (pp *pitchPack) accidental() rune {
	switch pp.halfStep {
	case -1:
		return 'b'
	case 1:
		return '#'
	default:
		return 0
	}
}
