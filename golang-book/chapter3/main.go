package main

import "fmt"

func main() {
	feetToMeter()
	ftoc()
}

func ftoc() {
	fmt.Print("Enter the temp in F: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * 5 / 9

	fmt.Println(input, "F equas", celsius, "C")
}

func feetToMeter() {
	fmt.Println("This is a program will convert feet into meters!")
	fmt.Print("Enter the feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048

	fmt.Println(input, "feet equas", meters, "meters")
}
