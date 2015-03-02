/*	Author: Henry Sarabia
	Project: Golang Book Homework
	Completion: 70%
 */

package main

import (
	"fmt"
	"strings"
)

func promptUser() string {
	var input1 string
	fmt.Print("This program can perform several tasks. Please enter one of the following to continue:\n")
	fmt.Print("Temperature_Converter \nDistance_Converter \nDiv_By_Three \nFizz_Buzz \nFind_Smallest \n")
	fmt.Print("Half_Int_Bool \nOdd_Number_Generator \nFibonacci \nInt_Swap \nFind_Perimeter \nSleep_Function \n")

	fmt.Scan(&input1)

	return input1
}
func tempPrompt() {
	var temp float32
	fmt.Print("Please enter the temperature in Fahrenheit: ")

	fmt.Scan(&temp)

	temp = (((temp - 32) * (5)) / 9)

	fmt.Printf("Converted temperature: %.2f Celsius", temp)
}

func distPrompt() {
	var dist float32
	fmt.Print("Please enter the distance in feet: ")

	fmt.Scan(&dist)

	dist = dist * 0.3048

	fmt.Printf("Converted distance: %.2fm", dist)
}

func divPrompt() {
	for i := 3; i <= 99; i = i + 3 {
		fmt.Println(i)
	}
}

func fizzPrompt() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func smallPrompt() {
	x := []int{ 48, 96, 86, 68, 57, 82, 63,70, 37, 34, 83, 27, 19, 97, 9, 17,}
	lowest := x[0]
	for _, value := range x {
		if value < lowest {
			lowest = value
		}
	}

	fmt.Print("The lowest value is: ", lowest)
}

func halfPrompt() (int, bool) {
	var input int
	fmt.Print("Please enter an integer: ")

	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Print("(", input, ", True)\n")
		return input, true
	} else {
		fmt.Print("(", input, ", False)\n")
		return input, false
	}
}

func largePrompt() {
	input := []int{15, 84, 32, 98, 5, 74, 71, 21, 64, 25,}

	output := findGreatest(input...)

	fmt.Print("The largest value is: ", output)
}

func findGreatest(nums ...int) int {
	largest := nums[0]
	for _, value := range nums {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func oddPrompt() {
	i := 1
	makeOdd := func() (output int) {
		output = i
		i += 2
		return
	}
	var input int
	fmt.Print("How many odd numbers would you like: ")
	fmt.Scan(&input)
	for j := 0; j < input; j++ {
		fmt.Println(makeOdd())
	}
}

func fibPrompt() {

}

func swapPrompt() {

}

func periPrompt() {

}

func sleepPrompt() {

}

func main() {
	choice := promptUser()

	if strings.EqualFold(choice, "Temperature_Converter") {
		tempPrompt()
	} else if strings.EqualFold(choice, "Distance_Converter") {
		distPrompt()
	} else if strings.EqualFold(choice, "Div_By_Three") {
		divPrompt()
	} else if strings.EqualFold(choice, "Fizz_Buzz") {
		fizzPrompt()
	} else if strings.EqualFold(choice, "Find_Smallest") {
		smallPrompt()
	} else if strings.EqualFold(choice, "Half_Int_Bool") {
		halfPrompt()
	} else if strings.EqualFold(choice, "Find_Largest") {
		largePrompt()
	} else if strings.EqualFold(choice, "Odd_Number_Generator") {
		oddPrompt()
	} else if strings.EqualFold(choice, "Fibonacci") {
		fibPrompt()
	} else if strings.EqualFold(choice, "Int_Swap") {
		swapPrompt()
	} else if strings.EqualFold(choice, "Find_Perimeter") {
		periPrompt()
	} else if strings.EqualFold(choice, "Sleep_Function") {
		sleepPrompt()
	}
}
