package main

import "fmt"

func nameInput() string {
	var name string

	fmt.Println("Enter your name")
	fmt.Scanln(&name)

	fmt.Println("Hello", name)
	return name

}

// func summation() int {

// 	var num1 int
// 	var num2 int

// 	fmt.Println("Enter first number")
// 	fmt.Scanln(&num1)

// 	fmt.Println("Enter second number")
// 	fmt.Scanln(&num2)

// 	sum := num1 + num2

// 	fmt.Println("sum", sum)
// 	return sum
// }

func getNum() (int, int) {
	var num1 int
	var num2 int

	fmt.Println("Enter first number")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number")
	fmt.Scanln(&num2)
	return num1, num2

}

func add(num1 int, num2 int) int {

	sum := num1 + num2
	return sum
}

func main() {

	// var name string

	// fmt.Println("Enter your name")
	// fmt.Scanln(&name)

	// fmt.Println("Hello", name)

	nameInput()

	// summation()

	p, q := getNum()

	sum := add(p, q)
	fmt.Println("This is summation", sum)

	// var num1 int
	// var num2 int

	// fmt.Println("Enter first number")
	// fmt.Scanln(&num1)

	// fmt.Println("Enter second number")
	// fmt.Scanln(&num2)

	// sum := num1 + num2

	// fmt.Println("sum", sum)
}
