// TODO: make it able to parse whatever I type

package main

import (
	"fmt"
)

func main() {
	prompt()
	input()
}

func prompt() {
	fmt.Println("This is a basic calculator:")
	fmt.Println("The program will compute with the following: \"first_num\" \"operator\" \"second_num\"")
	fmt.Println("For example: 1 + 2, 3 / 4, 5 ^ 6")
	fmt.Println("The following operations are valid:")
	fmt.Println("+ - * / ^\n")
}

/*
func input() {
	var first int
	var sec int
	var op rune

	fmt.Print("First Number: ")
	fmt.Scan(&first)
	fmt.Print("Second Number: ")

	fmt.Scan(&sec)
	fmt.Print("Operator: ")
	fmt.Scanf("%c", &op)

	result := compute(first, sec, op)

	fmt.Printf("%v %c %v = %v", first, op, sec, result)


}
*/

func compute(first int, sec int, op rune) int {
	switch op {
	case '+':
		return first + sec
	case '-':
		return first - sec
	case '*':
		return first * sec
	case '/':
		return first / sec
	case '^':
		result := 1
		for range sec {
			result *= first
		}
		return result
	}

	return 0
}
