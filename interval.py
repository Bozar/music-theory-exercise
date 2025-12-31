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
	if answer.isdigit():
		answer = int(answer)
	else:
		answer = -1

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
	check = right[1] - left[1] + 1
	#print(f'{answer}: {check}')
	return answer == check
