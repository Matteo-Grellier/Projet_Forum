package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Posts_Activity(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/posts_activity.html", "templates/layouts/sidebar.html")
	if err != nil {
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[SERVER_INFO_PAGE] : 🟠 Template execution: %s", err)
		return
	}
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "[SERVER_INFO_PAGE] : 🟢 Page 'posts_activity'")
	t.Execute(w, nil)
}
