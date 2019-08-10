package sum

//SumIndividual calculates sum of each elements in an giving slides
func SumIndividual(numbers []int) (sum int) {
	for _, val := range numbers {
		sum += val
	}
	return
}

//SumAll is a variadic function, which calculates sum of elements of each giving slides and return a slides of results
func SumAll(numbersToSum ...[]int) (sum []int) {
	lenOfNumbers := len(numbersToSum)
	sum = make([]int, lenOfNumbers) //initialize size of sum, equal to amounts of function arguments

	for i, numbers := range numbersToSum {
		sum[i] = SumIndividual(numbers)
	}

	return
}

//SumAllTails is a little different from SumAll, the former excludes the first element of every slides
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tails := numbers[1:]
			sum = append(sum, SumIndividual(tails))
		}
	}

	return
}
