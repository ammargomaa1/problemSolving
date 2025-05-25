func dailyTemperatures(temperatures []int) []int {
    output := make([]int, len(temperatures))
	stack := [][]int{}

	for ind, temp := range temperatures{
		if len(temperatures) > ind + 1 && temp < temperatures[ind + 1] {
			output[ind] =  1
			for len(stack) > 0 {
				if stack[len(stack)-1][1] < temperatures[ind + 1] {
					output[stack[len(stack)-1][0]] = ind + 1 - stack[len(stack)-1][0]
					stack = stack[:len(stack) - 1]
				}else{
					break
				}
			}
			continue
		}

		stack = append(stack, []int{ind, temp})

	}

	return output
}