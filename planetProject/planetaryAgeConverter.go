//	Author: Henry Sarabia
//	Date: 2/8/14
//	SID: 107174372

//	Planetary Age Converter
//	This small program will take in a user's age and a planet of their choosing as input
// 	which will be used to calculate their age based on the orbit of the given planet.
//	The program can detect whether or not the user is inputting valid data and will continue
//	to prompt the user for data	until the input is valid or the program is terminated.

package main

import "fmt"
import "strings"

//	These constants are the factors used in calculating the different planetary years
//	compared to Earth years.
const (
	YEAR 		= 365
	MERCURY 	= 4.15
	VENUS		= 1.63
	EARTH 		= 1.00
	MARS 		= 0.53
	JUPITER 	= 0.08
	SATURN 		= 0.03
	URANUS 		= 0.012
	NEPTUNE 	= 0.006
	PLUTO 		= 0.004
)

//	This struct is the basic data type used in the program to store the user's input.
type Profile struct {
	age 	float32
	planet	string
}
//	This function will prompt the user with a welcome message briefly explaining what the program does then
//	asks for their age as input. The program will check to make sure that the age is at least 1 year until
//	it will execute. Once valid data has been entered, the function will prompt the user for the name of a
//	planet from our solar system and will not accept any other strings. One aspect of the function to note
// 	is that the string comparison is not case sensitive.
func Prompt(input1 *float32, input2 *string) {
	fmt.Println("This program will calculate how old you would be on any \nof the planets from our solar system.")
	fmt.Print("Please enter your age: ")

	fmt.Scan(input1)

	//	This is the simple for loop that checks for whether the data entered is at least 1. If it is not
	//	then the program will prompt the user again and repeats the loop until valid data is entered.
	for *input1 < 1 {
		fmt.Print("Please enter an age of at least 1: ")
		fmt.Scan(input1)
	}

	fmt.Print("Please enter the planet name: ")

	fmt.Scan(input2)

	//	This is the for loop that checks for whether the string entered matches the name of a planet
	//	from our solar system. If the string does not match, the for loop will repeat until the data is valid.
	//	The string comparison function used is especially useful because it is case insensitive.
	for !strings.EqualFold(*input2, "Mercury") && !strings.EqualFold(*input2, "Venus") && !strings.EqualFold(*input2, "Earth") && !strings.EqualFold(*input2, "Mars") && !strings.EqualFold(*input2, "Jupiter") && !strings.EqualFold(*input2, "Saturn") && !strings.EqualFold(*input2, "Uranus") && !strings.EqualFold(*input2, "Neptune") && !strings.EqualFold(*input2, "Pluto") {
		fmt.Print("Please enter one of the 9 planets in our solar system: ")
		fmt.Scan(input2)
	}
	}

//	This is the function that converts the user's age input into the desired planetary age. The function takes in
//	the user.age and user.planet variables as its arguments and outputs two different float32 variables -
//	the new calculated age and the number of earth days that would equal, respectively.
func Convert(years float32, planet string) (float32, float32) {

	//	The following commented out code is the much more concise but error prone version of string comparison.
/*
switch planet {
		case "Mercury": return years * MERCURY, years * MERCURY * YEAR
		case "Venus": return years * VENUS, years * VENUS * YEAR
		case "Earth": return years, years * YEAR
		case "Mars": return years * MARS, years * MARS * YEAR
		case "Jupiter": return years * JUPITER, years * JUPITER * YEAR
		case "Saturn" : return years * SATURN, years * SATURN * YEAR
		case "Uranus": return years * URANUS, years * URANUS * YEAR
		case "Neptune": return years * NEPTUNE, years * NEPTUNE * YEAR
		case "Pluto": return years * PLUTO, years * PLUTO * YEAR
		default: return 0, 0
	}
*/

	//	This section of code matches the user.planet variable to one of the 9 planets in order to calculate the
	//	user's age on the corresponding planet by multiplying the user.age variable by the previously mentioned
	//	factor constants. The string comparison is done through the strings.EqualFold(s, t, string) function that
	//	returns a bool depending on whether or not the strings match without case sensitivity.
	if strings.EqualFold(planet, "Mercury") {
		return years * MERCURY, years * MERCURY * YEAR
	} else if strings.EqualFold(planet, "Venus") {
		return years * VENUS, years * VENUS * YEAR
	} else if strings.EqualFold(planet, "Earth") {
		return years, years * YEAR
	} else if strings.EqualFold(planet, "Mars") {
		return years * MARS, years * MARS * YEAR
	} else if strings.EqualFold(planet, "Jupiter") {
		return years * JUPITER, years * JUPITER * YEAR
	} else if strings.EqualFold(planet, "Saturn") {
		return years * SATURN, years * SATURN * YEAR
	} else if strings.EqualFold(planet, "Uranus") {
		return years * URANUS, years * URANUS * YEAR
	} else if strings.EqualFold(planet, "Neptune") {
		return years * URANUS, years * URANUS * YEAR
	} else if strings.EqualFold(planet, "Pluto") {
		return years * PLUTO, years * PLUTO * YEAR
	} else {
		return 0, 0
	}

}

//	This is the main function that will call the previously created functions. Once the functions execute and
//	the data is calculated, the main function will format and output the results to the user.
func main() {
	user := Profile{}
	Prompt(&user.age, &user.planet)
	newAge, newDays := Convert(user.age, user.planet)

	fmt.Println("Your age is: ", user.age)
	fmt.Println("Planet: ", user.planet)

	fmt.Print("Your age on ", user.planet)
	fmt.Printf(" is %.2f years.\n", newAge)

	fmt.Printf("That's %.f Earth days!\n", newDays)
}
