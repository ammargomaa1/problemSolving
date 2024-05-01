func dailyTemperatures(temperatures []int) []int {
    lenResult := len(temperatures)
	result := make([]int, lenResult)
	stack := [][]int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > stack[len(stack)-1][0] {
			_, index := stack[len(stack)-1][0], stack[len(stack)-1][1]
			result[index] = i - index
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, []int{temp, i})

	}

	return result
}