import interval as it


def main():
	count_right = 0
	left, right = 0, 0
	for i in range(0, it.MAX_PROBLEM):
		left, right = it.get_problem(i)
		count_right = it.verify_answer(count_right, left, right)
	it.get_result(count_right)


main()
