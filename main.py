import interval as it


def main():
	check = ''
	count_result = 0
	for i in range(0, it.MAX_PROBLEM):
		check = it.print_problem(i)
		count_result = it.verify_answer(count_result, check)
	it.print_result(count_result)


main()
