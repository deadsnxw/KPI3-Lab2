package lab2

import (
	"fmt"
	"io"
	"os"
)

type Computable interface {
	Compute() error
}

type ComputeHandler struct {
	Input  string
	Output string
}

type MockComputeHandler struct {
	Input  string
	Output string
}

func (ch *ComputeHandler) Compute() error {
	err := isValidExpression(ch.Input)
	if err != nil {
		return err
	}

	result := ConvertPostfixToPrefix(ch.Input)

	if ch.Output == "" {
		fmt.Println(result)
	} else {
		file, err := os.Create(ch.Output)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.WriteString(file, result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ch *MockComputeHandler) Compute() (string, error) {
	err := isValidExpression(ch.Input)
	if err != nil {
		return "", err
	}

	result := ConvertPostfixToPrefix(ch.Input)

	return result, nil
}
