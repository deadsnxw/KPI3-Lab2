package lab2

import (
	"errors"
	"strconv"
	"strings"
)

func isValidExpression(postfixExpression string) error {
	stack := []string{}

	for _, token := range strings.Split(postfixExpression, "") {
		if token == " " {
			continue
		}

		if isDigit(token) {
			stack = append(stack, token)
		} else {
			if len(stack) < 2 {
				return errors.New("incorrect expression")
			}

			stack = stack[:len(stack)-2]

			stack = append(stack, "result")
		}
	}

	if len(stack) != 1 {
		return errors.New("incorrect expression")
	}

	return nil
}

func ConvertPostfixToPrefix(postfixExpression string) string {
	stack := []string{}

	for _, token := range strings.Split(postfixExpression, "") {
		if token == " " {
			continue
		}
		if isDigit(token) {
			stack = append(stack, token)
		} else {
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var expression string
			if (token == "+" || token == "-") && len(operand1) != 1 {
				expression = token + " " + operand2 + " " + operand1
			} else {
				expression = token + " " + operand1 + " " + operand2
			}

			stack = append(stack, expression)
		}
	}

	return stack[0]
}

func isDigit(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}
