package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	User_pseudo   string
	User_password string
}

func OpenDataBase() *sql.DB {
	/*Ouverture de la base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if err != nil {
		fmt.Println("Could Not open Database")
	}
	return db
}
func GetLogin(w http.ResponseWriter, r *http.Request) {

	db := OpenDataBase()
	var eachPseudo User
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
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
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for verifPseudo.Next() {
		verifPseudo.Scan(&eachPseudo.User_pseudo)
		if eachPseudo.User_pseudo == pseudo {
			pseudoFound = true
			break
		}
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
	if passwordFound && pseudoFound {
		fmt.Println("VOUS ETES CONNECTER")
	} else {
		fmt.Println("VOUS NETES PAS CONNECTER")
	}

	fmt.Println(pseudo, password)
	http.Redirect(w, r, "/all_categories", http.StatusSeeOther)
}
