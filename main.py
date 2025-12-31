import interval as it


def main():
	count_right = 0
	for i in range(0, it.MAX_PROBLEM):
		it.get_problem(i)
		count_right = it.verify_answer(count_right)
	it.get_result(count_right)


main()
