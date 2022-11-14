package main

import "fmt"

func addition(a, b int) int { //addition of two numbers

	return a + b
}

func main() {

	//Create two variables
	var numberOne, numberTwo int

	//Take input from user
	fmt.Print("Enter first number : ")
	fmt.Scanln(&numberOne)
	fmt.Print("Enter second number : ")
	fmt.Scanln(&numberTwo)

	//Print addition of two numbers
	fmt.Println(addition(numberOne, numberTwo))
}
