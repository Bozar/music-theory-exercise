import interval as it


def main():
	check = ''
	count_result = 0
	is_ok = True
	for i in range(0, it.MAX_PROBLEM):
		check = it.print_problem(i)
		count_result, is_ok = it.verify_answer(count_result, check)
		if not is_ok:
			break
	it.print_result(count_result)


main()
