package stackque

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		tokenss := tokens[i]
		if tokenss == "+" || tokenss == "-" || tokenss == "*" || tokenss == "/" {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			c := 0
			switch tokenss {
			case "+":
				c = a + b
			case "-":
				c = b - a
			case "*":
				c = a * b
			case "/":
				c = b / a
			}
			stack = append(stack, c)
		} else {
			a, _ := strconv.Atoi(tokenss)
			stack = append(stack, a)
		}
	}
	return stack[0]
}
