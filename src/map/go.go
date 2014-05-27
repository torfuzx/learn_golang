package main

import (
	"fmt"
)

type Point struct {x, y, z int}
func (point Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}

func main() {
	{
		massFotPlanet := make(map[string]float64) // same as map[string]float{}
		massFotPlanet["Mercury"] 	= 0.06
		massFotPlanet["Venus"] 		= 0.82
		massFotPlanet["Earth"] 		= 1.00
		massFotPlanet["Mars"] 		= 0.11
		fmt.Println(massFotPlanet)
	}

	{
		triangle := make(map[*Point]string, 3)
		triangle[&Point{89, 47, 27}] = "α"
		triangle[&Point{86, 65, 86}] = "β"
		triangle[&Point{7, 44, 45}]  =  "γ"
		fmt.Println(triangle)
	}

	{
		nameForPoint := map[Point]string{} // same as make([Point]string)
		nameForPoint[Point{54, 91, 78}] = "x"
		nameForPoint[Point{54, 158, 89}] = "y"
		fmt.Println(nameForPoint)
	}

	{
		populationForCity := map[string]int {
			"Istanbul" : 10620000,
			"Mumbai": 12690000,
			"Shanghai": 13680000,
		}
		for city, population := range populationForCity {
			fmt.Printf("%-10s %8d\n", city, population)
		}

		// query the map
		population := populationForCity["Mumbai"]
		fmt.Println("Mumbai's population is ", population)

		population = populationForCity["Emerald City"]
		fmt.Println("Emerald City's population is ", population)
	}

	{
		populationForCity := map[string]int {
			"Istanbul" : 10620000,
			"Mumbai": 12690000,
			"Shanghai": 13680000,
		}

		city := "Istanbul"
		if population, found := populationForCity[city]; found {
			fmt.Printf("%a's population is %d\n", city, population)
		} else {
			fmt.Printf("%a's population data is unavailable\n", city)
		}

		city = "Emerald City"
		_, present := populationForCity[city]
		fmt.Printf("%q is in the map == %t\n", city, present)
	}
}



