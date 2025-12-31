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


def print_problem(p_index):
	left_index = rd.randrange(0, MAX_NOTE_INDEX)
	right_index = left_index
	while right_index == left_index:
		right_index = rd.randrange(0, MAX_NOTE_INDEX)
	left = NOTE_DATA[left_index]
	right = NOTE_DATA[right_index]
	print(f'{p_index + 1:02}: {left[0]} - {right[0]}')
	return left_index, right_index


def verify_answer(count_right, left_index, right_index):
	answer = input('$?: ')
	#if answer.isdigit():
		#answer = int(answer)
	#else:
		#answer = -1
	if not _is_right_answer(answer, left_index, right_index):
		print('Wrong answer')
	else:
		count_right += 1
	return count_right


def print_result(count_right):
	print(f'$$: {count_right}/{MAX_PROBLEM}')


def _is_right_answer(answer, left_index, right_index):
	if left_index == right_index:
		raise ValueError('Left == Right: {left}')
		return False

	left = NOTE_DATA[left_index]
	right = NOTE_DATA[right_index]
	if left_index > right_index:
		right = NOTE_DATA[right_index + MAX_NOTE_INDEX]
	size = _get_size(left[1], right[1])
	quality = _get_quality(size, left[2], right[2])
	check = f'{quality}{size}'
	#print(f'{answer}: {check}')
	return answer == check


def _get_size(left_size, right_size):
	return right_size - left_size + 1


def _get_quality(size, left_step, right_step):
	step = right_step - left_step
	quality = STEP_TO_QUALITY.get(f'{size}-{step}', '')
	if quality == '':
		raise ValueError(f'No quality: {size}-{left_step}:{right_step}')
	return quality
