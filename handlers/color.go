package handlers

import (
	"fmt"
)

// Affichage des logs en couleur
func Color(couleur int, str string) {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	colorYellow := "\033[33m"
	if couleur == 1 {
		fmt.Println(string(colorGreen), str)
	} else if couleur == 2 {
		fmt.Println(string(colorBlue), str)
	} else if couleur == 3 {
		fmt.Println(string(colorYellow), str)
	} else {
		fmt.Println(string(colorRed), str)
	}
}
