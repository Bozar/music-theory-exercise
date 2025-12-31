#import


MAX_PROBLEM = 10


def get_problem(p_index):
	print(f'Problem {p_index}: ')


def verify_answer(count_right):
	answer = input('?: ')
	if answer.isdigit():
		answer = int(answer)
	else:
		answer = -1

	if not _is_right(answer):
		print('Wrong answer')
	else:
		count_right += 1
	return count_right


def get_result(count_right):
	print(f'{count_right}/{MAX_PROBLEM}')


def _is_right(answer):
	if answer % 2 == 0:
		return True
	return False
