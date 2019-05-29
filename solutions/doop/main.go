package main

import (
	"fmt"
	"os"
	"strconv"
)

func validOperator(test string) bool {
	operatorsTable := []string{"+", "-", "*", "/", "%"}
	result := false
	for _, operator := range operatorsTable {
		if operator == test {
			result = true
		}
	}
	return result
}

func main() {

	if len(os.Args) == 4 {
		if validOperator(os.Args[2]) == false {
			fmt.Println(0)
			return
		}
		var rest int
		firstNumber, _ := strconv.Atoi(os.Args[1])
		secondNumber, _ := strconv.Atoi(os.Args[3])

		if secondNumber == 0 && os.Args[2] == "/" {
			fmt.Print("No division by 0\n")
			return
		} else if secondNumber == 0 && os.Args[2] == "%" {
			fmt.Print("No Modulo by 0\n")
			return
		} else if os.Args[2] == "+" {
			rest = firstNumber + secondNumber
		} else if os.Args[2] == "-" {
			rest = firstNumber - secondNumber
		} else if os.Args[2] == "/" {
			rest = firstNumber / secondNumber
		} else if os.Args[2] == "*" {
			rest = firstNumber * secondNumber
		} else if os.Args[2] == "%" {
			rest = firstNumber % secondNumber
		}
		fmt.Println(rest)
	} else {
		return
	}
}
