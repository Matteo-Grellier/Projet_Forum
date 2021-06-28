package BDD

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Vérifie si l'élément demandé est déjà dans la base de données
func VerifyBDD(element string, column string) (bool, string) {
	db := OpenDataBase()
	var oneElement string
	var prepareElements *sql.Stmt
	var errorPrepare error
	if column == "pseudo" {
		prepareElements, errorPrepare = db.Prepare("SELECT pseudo FROM user")
	} else if column == "mail" {
		prepareElements, errorPrepare = db.Prepare("SELECT mail FROM user")
	} else if column == "session" {
		prepareElements, errorPrepare = db.Prepare("SELECT user_pseudo FROM session")
	} else if column == "like" {
		prepareElements, errorPrepare = db.Prepare("SELECT user_pseudo FROM like")
	}

	if errorPrepare != nil {
		log.Fatal(errorPrepare)
	}

	allElements, err := prepareElements.Query()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for allElements.Next() {
		allElements.Scan(&oneElement)
		if oneElement == element {
			ErrorMessage := column + " déjà dans la base de données."
			allElements.Close()
			return true, ErrorMessage
		}
	}
	allElements.Close()
	return false, ""
}

// vérification de la correspondance entre le mdp de l'utilisateur et le mdp entré
func VerifyPassword(password string, pseudo string) bool {
	db := OpenDataBase()
	var eachPseudo User
	selectPassword, _ := db.Prepare("SELECT password FROM user WHERE pseudo = ?")
	verifPassword, err := selectPassword.Query(pseudo)
	if err != nil {
		log.Fatal(err)
	}
	for verifPassword.Next() {
		verifPassword.Scan(&eachPseudo.Password)
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.Password), []byte(password))
		if err != nil {
			log.Println(err)
			verifPassword.Close()
			return false
		} else {
			verifPassword.Close()
			return true
		}
	}
	db.Close()
	return false
}
