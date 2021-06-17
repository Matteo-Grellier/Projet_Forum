package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Println("Page Inscription ✅")
	VerifyCookie(w, r)
	t.Execute(w, nil)
}
