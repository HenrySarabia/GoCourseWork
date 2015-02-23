package main

import "fmt"

func promptUser(input1 *float32) {
	fmt.Print("Please enter the temperature in Fahrenheit: ")

	fmt.Scan(input1)
}

func main() {
	var temp float32

	promptUser(&temp)

	temp = (((temp - 32) * (5)) / 9)

	fmt.Printf("Converted temperature: %.2f Celsius", temp)
}
