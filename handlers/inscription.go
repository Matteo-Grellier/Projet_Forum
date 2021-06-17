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
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[SERVER_INFO_PAGE] : 🟠 Template execution: %s", err)
		return
	}

	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "[SERVER_INFO_PAGE] : 🟢 Page 'inscription'")
	t.Execute(w, nil)
}

/* func getPassword(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var e string = r.PostFormValue("Password")
	fmt.Println(string("\033[1;37m\033[0m"), e)
}
*/
