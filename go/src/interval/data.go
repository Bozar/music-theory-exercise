package interval

const maxProblem int = 10
const maxPitchPacksIndex int = 7

const minStep int = -1
const maxStep int = 2

var pitchPacks = []pitchPack{
	pitchPack{'C', 0, 0, 0},
	pitchPack{'D', 1, 0, 2},
	pitchPack{'E', 2, 0, 4},
	pitchPack{'F', 3, 0, 5},
	pitchPack{'G', 4, 0, 7},
	pitchPack{'A', 5, 0, 9},
	pitchPack{'B', 6, 0, 11},
	pitchPack{'C', 7, 0, 12},
	pitchPack{'D', 8, 0, 14},
	pitchPack{'E', 9, 0, 16},
	pitchPack{'F', 10, 0, 17},
	pitchPack{'G', 11, 0, 19},
	pitchPack{'A', 12, 0, 21},
	pitchPack{'B', 13, 0, 23},
}

var stepToQuality = map[int]rune{
	202: 'M',
	201: 'm',
	203: 'A',
	200: 'd',

	304: 'M',
	303: 'm',
	305: 'A',
	302: 'd',

	405: 'P',
	406: 'A',
	404: 'd',

	507: 'P',
	508: 'A',
	506: 'd',

	609: 'M',
	608: 'm',
	610: 'A',
	607: 'd',

	711: 'M',
	710: 'm',
	712: 'A',
	709: 'd',
}
