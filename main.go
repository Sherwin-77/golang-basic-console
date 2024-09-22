package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read input as uint8
func readUint8(input *bufio.Reader, inputMessage string) uint8 {
	var rawStr string
	var n uint8

	fmt.Print(inputMessage)
	rawStr, _ = input.ReadString('\n')
	tmp, err := strconv.ParseUint(strings.TrimSpace(rawStr), 10, 8)
	n = uint8(tmp)
	// If error, print error message and ask for input again
	if err != nil {
		fmt.Println("Invalid input. Please input a valid number")
		return readUint8(input, inputMessage)
	}

	return n
}

// Read input as uint32
func readUint32(input *bufio.Reader, inputMessage string) uint32 {
	var rawStr string
	var n uint32

	fmt.Print(inputMessage)
	rawStr, _ = input.ReadString('\n')
	tmp, err := strconv.ParseUint(strings.TrimSpace(rawStr), 10, 32)
	n = uint32(tmp)
	// If error, print error message and ask for input again
	if err != nil {
		fmt.Println("Invalid input. Please input a valid number")
		return readUint32(input, inputMessage)
	}

	return n
}

// Read input as float32
func readFloat32(input *bufio.Reader, inputMessage string) float32 {
	var rawStr string
	var n float32

	fmt.Print(inputMessage)
	rawStr, _ = input.ReadString('\n')
	tmp, err := strconv.ParseFloat(strings.TrimSpace(rawStr), 32)
	n = float32(tmp)
	// If error, print error message and ask for input again
	if err != nil {
		fmt.Println("Invalid input. Please input a valid number")
		return readFloat32(input, inputMessage)
	}

	return n
}

// Variadic usage to echo input with "UwU" appended
func uwuEcho(args ...string) {
	fmt.Print("UwU ")
	for _, arg := range args {
		fmt.Print(arg, " UwU ")
	}
	fmt.Println()
}

func mathOperation(input *bufio.Reader, mathHistory *[]string) {
	var firstNumber float32 = readFloat32(input, "Input first number: ")
	var secondNumber float32 = readFloat32(input, "Input second number: ")

	fmt.Print("Input operator (+, -, *, /): ")
	var operator string
	var result float32
	// Loop until valid operator is inputted
	for {
		operator, _ = input.ReadString('\n')
		operator = strings.TrimSpace(operator)

		switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			result = firstNumber / secondNumber
		default:
			fmt.Print("Invalid operator. Please input a valid operator (+, -, *, /):")
			continue
		}
		break
	}

	// Add math operation to history and print result. %g to print float without trailing zeros
	*mathHistory = append(*mathHistory, fmt.Sprintf("%g %s %g = %g", firstNumber, operator, secondNumber, result))
	fmt.Println("Result:", result)
}

func showMathHistory(mathHistory []string) {
	var operators = [4]string{"+", "-", "*", "/"}

	// Group math history by operator
	var groupedHistory = func() map[string][]string {
		grouped := make(map[string][]string)
		for _, entry := range mathHistory {
			for _, operator := range operators {
				if strings.Contains(entry, operator) {
					grouped[operator] = append(grouped[operator], entry)
					break
				}
			}
		}
		return grouped
	}()

	// Print history for each operator
	for operator, history := range groupedHistory {
		fmt.Println("History for operator", operator)
		for i, entry := range history {
			if i > 9 {
				break
			}
			fmt.Println(entry)
		}
	}
}

func clearMathHistory(mathHistory *[]string) {
	*mathHistory = []string{}
	fmt.Println("Math history cleared")
}

// Recursive fast fibonacci algorithm. See https://en.wikipedia.org/wiki/Fibonacci_number#Matrix_form
func fastFibonacci(n uint32) (uint32, uint32) {
	if n == 0 {
		return 0, 1
	}
	var a, b uint32 = fastFibonacci(n / 2)
	var c uint32 = a * (b*2 - a)
	var d uint32 = a*a + b*b
	if n%2 == 0 {
		return c, d
	} else {
		return d, c + d
	}
}

func countFibonacci(input *bufio.Reader) {
	var n uint32 = readUint32(input, "Input n: ")
	var result, _ = fastFibonacci(n)

	fmt.Println("Fibonacci number at index", n+1, "is", result)
}

func main() {
	const mainMenu = `
	1. Show Hello world
	2. UwUify
	3. Math operation
	4. Show math history
	5. Clear math history
	6. Count fibonacci
	7. Exit
	`

	var input *bufio.Reader = bufio.NewReader(os.Stdin)
	var mathHistory []string
	for {
		fmt.Println(mainMenu)
		fmt.Println("╭─Main Menu")
		var option uint8 = readUint8(bufio.NewReader(input), "╰─Choose menu: ")

		switch option {
		case 1:
			fmt.Println("Hello World")
		case 2:
			fmt.Print("Input text: ")
			var line, err = input.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input")
				continue
			}
			uwuEcho(strings.Split(strings.TrimSpace(line), " ")...)
		case 3:
			mathOperation(input, &mathHistory)
		case 4:
			showMathHistory(mathHistory)
		case 5:
			clearMathHistory(&mathHistory)
		case 6:
			countFibonacci(input)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid menu")
		}
	}
}
