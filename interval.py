import random as rd


MAX_PROBLEM = 10
PROBLEMS = [
	'C', 'D', 'E', 'F', 'G', 'A', 'B',
]


def get_problem(p_index):
	max_index = len(PROBLEMS) - 1
	left = rd.randint(0, max_index)
	right = left
	while right == left:
		right = rd.randint(0, max_index)
	print(f'{p_index + 1}: {PROBLEMS[left]} - {PROBLEMS[right]}')
	return left, right


def verify_answer(count_right, left, right):
	answer = input('?: ')
	if answer.isdigit():
		answer = int(answer)
	else:
		answer = -1

	if not _is_right_answer(answer, left, right):
		print('Wrong answer')
	else:
		count_right += 1
	return count_right


def get_result(count_right):
	print(f'{count_right}/{MAX_PROBLEM}')


def _is_right_answer(answer, left, right):
	check = 0
	if left < right:
		check = right - left + 1
		#print(f'{answer}: {check}')
		return answer == check
	elif left > right:
		check = 9 - (left - right + 1)
		#print(f'{answer}: {check}')
		return answer == check
	raise ValueError('Left == Right: {left}')
	return False
