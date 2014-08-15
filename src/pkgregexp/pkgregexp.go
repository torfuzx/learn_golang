package main

import (
	"regexp"
	"fmt"
)

func main() {
	{
		names := []string {
			// English Examples
			"Elizabeth Blackshaw",
			"Carl William Brandt",
			"Matthew William Harman Jr.",
			"Brian David Oviatte III",
			// Spanish Examples
			"Jose Estebán Martinez Delgado",
			"María Anna Padilla Lopez",
			"Juan Angel de la Cruz Vasquez Ovalle",
			// Portuguese Examples
			"Bernardo Duarte Preto",
			"Maria Angelica Domingues Lima",
			"Maria da Conceição",
			// French Examples
			"René Emile Donné",
			"Charlette Antoinette Vuatrin",
			"Anne Gerard-George",
			"Jean-Baptiste Boulanger",
			// German Examples
			"Anna Katharine Dorothea Weiß",
			"Carl Friedrich Wilhelm Müller",
			"Adolph Mein",
			// Scandinavian Examples
			"Thor Mikkel Hegland",
			"Kristi Olavsdatter Lien",
			"Mette Christine Gottfredson",
			"Kirstine Marie Ehlers",
			"Jörgen Christian Jensen",
			"Ingrid Hendricksdotter",
			"Jens Pederson",
			"Åsa Maria Kahkipuro",
			"Jaakko Maenpaa",
		}
		// first run
		nameRx := regexp.MustCompile(`(\pL+\.?(?:\s+\pL+\.?)*)\s+(\pL+)`)
		for i := 0; i < len(names); i++ {
			names[i] = nameRx.ReplaceAllString(names[i], "${2}, ${1}")
		}
		fmt.Printf("%#v\n", names)

		// second run - using named group
		nameRx = regexp.MustCompile(
			`(?P<fornames>\pL+\.?(?:\s+\pL+\.?)*)\s+(?P<surname>\pL+)`)
		for i := 0; i < len(names); i++ {
			names[i] = nameRx.ReplaceAllString(names[i], "${surname}, ${fornames}")
		}
		fmt.Printf("%#v\n", names)

	}
}

