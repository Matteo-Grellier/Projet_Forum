package handlers

import (
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

var UUID string
var PseudoConnected string

func GetLogin(w http.ResponseWriter, r *http.Request) {

	db := BDD.OpenDataBase()
	var eachPseudo User
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("templates/connection.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}

	pseudo := r.FormValue("Pseudo")
	pseudoFound := false
	password := r.FormValue("Password")
	passwordFound := false
	verifPseudo, err := db.Query("SELECT pseudo FROM user")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻  Could not query database")
		log.Fatal(err)
	}
	selectPassword, err := db.Prepare("SELECT password FROM user WHERE pseudo = ?")
	verifPassword, _ := selectPassword.Query(pseudo)
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for verifPseudo.Next() {
		verifPseudo.Scan(&eachPseudo.User_pseudo)
		if eachPseudo.User_pseudo == pseudo {
			pseudoFound = true
			break
		}
	}
	verifPseudo.Close()

	message := false
	if !pseudoFound {

		ErrorMessage = "Ce pseudonyme n'existe pas dans notre base de donnée"
	} else {
		message = true
	}
	for verifPassword.Next() {
		verifPassword.Scan(&eachPseudo.User_password)
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.User_password), []byte(password))
		if err != nil {
			log.Println(err)
		} else {
			passwordFound = true
		}
	}
	verifPassword.Close()
	if !passwordFound && message {
		ErrorMessage = "Ce mot de passe n'existe pas"
	}

	if passwordFound && pseudoFound {
		UUID = CreateCookie(w, r)
		CreateUUID(pseudo, UUID, db)
		Color(1, "[CONNEXION] : 🟢 Vous êtes connecté ")

		http.Redirect(w, r, "/all_categories", http.StatusSeeOther)
	} else {
		Color(4, "[CONNEXION] : 🔻 Vous n'êtes pas connecté ")

		ErrorsConnections := Errors{
			Error:  ErrorMessage,
			Pseudo: pseudo,
		}
		ErrorMessage = ""
		t.Execute(w, ErrorsConnections)
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")

}
