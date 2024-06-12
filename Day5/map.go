package main

import "fmt"

func main() {
	avengers := map[string]map[string]string{
		"Iron Man": {
			"profession": "Engineer",
			"speciality": "Powered Armor",
		},
		"Captain America": {
			"profession": "Soldier",
			"speciality": "Super Soldier Serum",
		},
		"Thor": {
			"profession": "God",
			"speciality": "God of Thunder",
		},
		"Hulk": {
			"profession": "Scientist",
			"speciality": "Gamma Radiation",
		},
		"Black Widow": {
			"profession": "Spy",
			"speciality": "Espionage",
		},
		"Hawkeye": {
			"profession": "Archer",
			"speciality": "Marksmanship",
		},
	}

	fmt.Println("Avenger character:")
	var input string
	fmt.Scanln(&input)

	if character, ok := avengers[input]; ok {
		fmt.Println("Profession:", character["profession"])
		fmt.Println("Specialty:", character["speciality"])
	} else {
		fmt.Println("Character not found.")
	}
}
