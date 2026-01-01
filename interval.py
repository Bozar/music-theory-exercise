import random as rd


MAX_PROBLEM = 10
MAX_NOTE_INDEX = 7

NOTE_DATA = (
	('C', 1, 0),
	('D', 2, 2),
	('E', 3, 4),
	('F', 4, 5),
	('G', 5, 7),
	('A', 6, 9),
	('B', 7, 11),
	('C', 8, 12),
	('D', 9, 14),
	('E', 10, 16),
	('F', 11, 17),
	('G', 12, 19),
	('A', 13, 21),
	('B', 14, 23),
)

STEP_TO_QUALITY = {
	'2-2': 'M',
	'2-1': 'm',
	'2-3': 'A',
	'2-0': 'd',

	'3-4': 'M',
	'3-3': 'm',
	'3-5': 'A',
	'3-2': 'd',

	'4-5': 'P',
	'4-6': 'A',
	'4-4': 'd',

	'5-7': 'P',
	'5-8': 'A',
	'5-6': 'd',

	'6-9': 'M',
	'6-8': 'm',
	'6-10': 'A',
	'6-7': 'd',

	'7-11': 'M',
	'7-10': 'm',
	'7-12': 'A',
	'7-9': 'd',
}

MIN_HALF_STEP = -1
MAX_HALF_STEP = 2
STEP_TO_SIGN = {
	-1: 'b',
	0: '',
	1: '#',
}


def print_problem(index_p):
	left, right, check = _get_problem()
	note_l, step_l = left
	note_r, step_r = right
	sign_l = STEP_TO_SIGN[step_l]
	sign_r = STEP_TO_SIGN[step_r]
	print(
			f'{index_p + 1:02}: '
			f'{note_l[0]}{sign_l} - '
			f'{note_r[0]}{sign_r}'
	)
	return check


def verify_answer(count_result, check):
	answer = input('$?: ')
	#if answer.isdigit():
		#answer = int(answer)
	#else:
		#answer = -1
	if answer == check:
		count_result += 1
	else:
		print('Wrong answer')
	#print(f'{answer}: {check}')
	return count_result


def print_result(count_result):
	print(f'$$: {count_result}/{MAX_PROBLEM}')


def _get_problem():
	while True:
		index_l = rd.randrange(0, MAX_NOTE_INDEX)
		index_r = index_l
		while index_r == index_l:
			index_r = rd.randrange(0, MAX_NOTE_INDEX)
		step_l = rd.randrange(MIN_HALF_STEP, MAX_HALF_STEP)
		step_r = rd.randrange(MIN_HALF_STEP, MAX_HALF_STEP)

		note_l = NOTE_DATA[index_l]
		note_r = NOTE_DATA[index_r]
		if index_l > index_r:
			note_r = NOTE_DATA[index_r + MAX_NOTE_INDEX]

		size = _get_size(note_l[1], note_r[1])
		quality = _get_quality(
				size, note_l[2] + step_l, note_r[2] + step_r
		)
		if quality != '':
			check = f'{quality}{size}'
			break
	return (note_l, step_l), (note_r, step_r), check


def _get_size(size_l, size_r):
	return size_r - size_l + 1


def _get_quality(size, step_l, step_r):
	step = step_r - step_l
	quality = STEP_TO_QUALITY.get(f'{size}-{step}', '')
	return quality
