func isValid(s string) bool {
    stack := []rune{}
    oppeningParentheses := map[string]string{
        "[" : "]",
        "(" : ")",
        "{" : "}",
    }

    for _, parenthes := range s {
        if _,ok := oppeningParentheses[string(parenthes)]; ok {
            stack = append(stack, parenthes)
            continue
        }

        if len(stack) > 0 && string(parenthes) == oppeningParentheses[string(stack[len(stack)-1])] {
            stack = stack[:len(stack)-1]
        }else{
            return false
        }

    }

    if len(stack) > 0 {
        return false
    }

    return true
}