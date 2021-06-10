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
	//var HashPass = hashPassword(password)
	verifPseudo, err := db.Query("SELECT pseudo FROM user")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	verifPassword, err := db.Query("SELECT password FROM user")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for verifPseudo.Next() {
		verifPseudo.Scan(&eachPseudo.User_pseudo)
		if eachPseudo.User_pseudo == pseudo {
			fmt.Println("LE PSEUDO EXISTE")
			pseudoFound = true
			break
		}
	}
	verifPassword.Scan(&eachPseudo.User_password)
	error := bcrypt.CompareHashAndPassword([]byte(eachPseudo.User_password), []byte(password))
	if error != nil {
		log.Println(err)
		return
	}
	passwordFound = true
	if passwordFound && pseudoFound {
		fmt.Println("VOUS ETES CONNECTER")
	}
	/* if pseudo != verifPseudo {
		fmt.Println("le pseudo n'est pas dans la BDD")
	} */
	/* var allDataLogin = []string{pseudo, password}
	verifyInput(allDataLogin) */

	fmt.Println(pseudo, password)
	http.Redirect(w, r, "/all_categories", http.StatusSeeOther)
}
