func generateParenthesis(n int) []string {
    arrSol = []string{}
    generate(n, "", 0)

    return arrSol
    
}

var arrSol []string

func generate(n int, sol string, pcounter int) {
	if n*2 == len(sol) {
		if pcounter == 0 && validateParentheses(sol) {
			arrSol = append(arrSol, sol)
		}
		return
	}

	generate(n, sol+"(", pcounter+1)
	if len(sol) != 0 {
		generate(n, sol+")", pcounter-1)
	}
	return
}

func validateParentheses(pa string) bool {
	stack := []rune{}

	for _, letter := range pa {

		if letter == '(' {
			stack = append(stack, letter)
		}

		if letter == ')' {
			if len(stack) == 0 {
				
				return false
			}

			stack = stack[:len(stack)-1]
		}

	}

	return true
}
