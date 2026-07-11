package main

import "fmt"

func getNum(num1 int, num2 int) (int, int) {

	sum := num1 + num2

	mul := num1 * num2

	return sum, mul

}

func text() {

	fmt.Println("Hello Everyone")
}

func hello(name string) {
	fmt.Println("hello", name)
}

func main() {

	hello("rafsan")
	text()

	a := 30
	b := 30

	p, q := getNum(a, b)

	fmt.Println(p)
	fmt.Println(q)

}
