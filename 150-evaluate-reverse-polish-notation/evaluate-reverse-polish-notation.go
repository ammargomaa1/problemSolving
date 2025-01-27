func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        tokenInt, err := strconv.Atoi(token)

        if err == nil {
            stack = append(stack, tokenInt)
        } else {
            num1 := stack[len(stack) - 1]
            num2 := stack[len(stack) - 2]
            stack = stack[:len(stack)-2]
			
			
            switch token {
                case "+":
                    stack = append(stack , num1 + num2)
                
                case "-":
                    stack = append(stack , num2 - num1)

                case "*":
                    stack = append(stack , num1 * num2)

                case "/":
                    stack = append(stack , num2 / num1)
            }

        }
    }

    return stack[0]
}