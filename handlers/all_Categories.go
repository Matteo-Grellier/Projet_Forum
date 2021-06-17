package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func All_Categories(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/all_categories.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[SERVER_INFO_PAGE] : 🟠 Template execution: %s", err)
		return
	}
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "[SERVER_INFO_PAGE] : 🟢 Page 'all_catégories'")
	t.Execute(w, nil)
}
