package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	User_pseudo   string
	User_password string
}

func GetLogin(w http.ResponseWriter, r *http.Request) {

	db := BDD.OpenDataBase()
	var eachPseudo User
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("templates/connection.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[SERVER_INFO_PAGE] : 🟠 Template execution: %s", err)
		return
	}

	pseudo := r.FormValue("Pseudo")
	pseudoFound := false
	password := r.FormValue("Password")
	passwordFound := false
	verifPseudo, err := db.Query("SELECT pseudo FROM user")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	verifPassword, err := db.Query("SELECT password FROM user WHERE pseudo = '" + pseudo + "'")
	if err != nil {
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[BDD_INFO] : 🔻 Template execution: %s", err)
	}
	for verifPseudo.Next() {
		verifPseudo.Scan(&eachPseudo.User_pseudo)
		if eachPseudo.User_pseudo == pseudo {
			pseudoFound = true
			break
		}
	}
	message := false
	if !pseudoFound {

		ErrorMessage = "Ce pseudonyme n'existe pas dans notre base de donnée"
	} else {
		message = true
	}
	for verifPassword.Next() {
		verifPassword.Scan(&eachPseudo.User_password)
		fmt.Println()
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.User_password), []byte(password))
		if err != nil {
			log.Println(err)
		} else {
			passwordFound = true
		}
	}
	if !passwordFound && message {
		ErrorMessage = "Ce mot de passe n'existe pas"
	}

	if passwordFound && pseudoFound {
		fmt.Println("VOUS ETES CONNECTÉ")
		fmt.Println(pseudo, password)

		http.Redirect(w, r, "/all_categories", http.StatusSeeOther)
	} else {
		fmt.Println("VOUS NETES PAS CONNECTER")
		ErrorsConnections := Errors{
			Error:  ErrorMessage,
			Pseudo: pseudo,
		}
		ErrorMessage = ""
		t.Execute(w, ErrorsConnections)
	}
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")

}
