package main

import (
	"fmt"
)

type Person struct {
	Title 		string
	Fornames 	[]string
	Surname		string
}

type Author1 struct  {
	Names 		Person
	Title 		[]string
	YearBorn 	int
}

type Author2 struct {
	Person
	Title []string
	YearBorn int
}

func main() {
	{
		points := [][2]int {{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
		for _, point := range points {
			fmt.Printf("(%d, %d)", point[0], point[1])
		}
		fmt.Println()
	}

	{
		points := []struct {x, y int} {{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
		for _, point := range points {
			fmt.Printf("(%d, %d)", point.x, point.y)
		}
		fmt.Println()
	}

	{
		author1 := Author1{
		Person{
			"Mr", []string{"Robert", "Louis", "Balfour"}, "Stevenson"},
			[]string{"Kiddnapped", "Treasure Island"},
			1850,
		}
		fmt.Printf("%#v\n", author1)

		author1.Names.Title = ""
		author1.Names.Fornames = []string{"Oscar", "Fingal", "O'Flaheritie", "Wills"}
		author1.Names.Surname = "Wilde"
		author1.Title = []string{"The picture of Dorian Gray"}
		author1.YearBorn += 4
		fmt.Printf("%#v\n", author1)
	}
}

