package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"

	lab2 "github.com/bifynok/KPI3-Lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")

	inputFilePath = flag.String("f", "", "Read expression in file to compute")

	outputFile = flag.String("o", "", "Output file")

	expression string
)

func main() {
	flag.Parse()

	if (*inputExpression != "" && *inputFilePath == "") || (*inputExpression == "" && *inputFilePath != "") {

		if *inputExpression != "" {
			expression = *inputExpression
		}

		if *inputFilePath != "" {
			file, err := os.Open(*inputFilePath)

			if err != nil {
				fmt.Println("Помилка відкриття файлу:", err)
				return
			}
			defer file.Close()

			buffer := make([]byte, 1024)

			for {
				n, err := file.Read(buffer)

				if err != nil && err != io.EOF {
					fmt.Println("Помилка читання з файлу:", err)
					return
				}

				if n == 0 {
					break
				}

				expression = string(buffer[:n])
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Incorrect using of atributes")
	}

	for _, char := range expression {
		if unicode.IsLetter(char) {
			fmt.Fprintln(os.Stderr, "Incorrect expression")
			os.Exit(1)
		}
	}

	handler := &lab2.ComputeHandler{}

	handler.Input = expression

	handler.Output = *outputFile

	err := handler.Compute()
	if err != nil {
		panic(err)
	}

}
