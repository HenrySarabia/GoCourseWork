package main

import (
	"fmt"
)

const(
	MERCURY  	= 5.43
	VENUS		= 5.20
	EARTH		= 5.51
	MARS		= 3.94
	JUPITER		= 1.33
	SATURN		= 0.69
	URANUS		= 1.27
	NEPTUNE		= 1.64
)

type Planet struct {
	Name string
	Density float64
}

func (p Planet) printDensity() {
	fmt.Printf("The density of %s is %.2f grams per cubic centimeter.\n", p.Name, p.Density)
}

func main() {

	planetSlice := []Planet{
		Planet{Name: "Mercury", Density: MERCURY, },
		Planet{Name: "Venus", Density: VENUS, },
		Planet{Name: "Earth", Density: EARTH, },
		Planet{Name: "Mars", Density: MARS, },
	}

	planetArray := [4]Planet{}
	planetArray[0] = Planet{Name: "Jupiter", Density: JUPITER, }
	planetArray[1] = Planet{Name: "Saturn", Density: SATURN, }
	planetArray[2] = Planet{Name: "Uranus", Density: URANUS, }
	planetArray[3] = Planet{Name: "Neptune", Density: NEPTUNE, }

	for i, _ := range planetArray {
		planetSlice = append(planetSlice, planetArray[i])
	}

	for _, current := range planetSlice {
		current.printDensity()
	}
}
