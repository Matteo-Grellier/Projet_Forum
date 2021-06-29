package BDD

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// Vérifie si l'élément demandé est déjà dans la base de données
func VerifyBDD(element string, column string) (bool, string, error) {
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
	}

	if errorPrepare != nil {
		db.Close()
		return false, "", errorPrepare
	}

	allElements, err := prepareElements.Query()
	if err != nil {
		db.Close()
		return false, "", err
	}
	for allElements.Next() {
		allElements.Scan(&oneElement)
		if oneElement == element {
			ErrorMessage := column + " déjà dans la base de données."
			allElements.Close()
			return true, ErrorMessage, nil
		}
	}
	allElements.Close()
	return false, "", nil
}

func VerifyLike(post_id int, user_pseudo string) (bool, int, error) {
	var oneUser string
	var statusLike int
	db := OpenDataBase()
	prepareLike, err := db.Prepare("SELECT user_pseudo, liked FROM like WHERE post_id = ?")
	if err != nil {
		db.Close()
		return false, 0, err
	}
	allLikes, err := prepareLike.Query(post_id)
	if err != nil {
		db.Close()
		return false, 0, err
	}
	for allLikes.Next() {
		allLikes.Scan(&oneUser, &statusLike)
		if oneUser == user_pseudo {
			allLikes.Close()
			return true, statusLike, nil
		}
	}
	allLikes.Close()
	return false, 0, nil
}

// vérification de la correspondance entre le mdp de l'utilisateur et le mdp entré
func VerifyPassword(password string, pseudo string) (bool, error) {
	db := OpenDataBase()
	var eachPseudo User
	selectPassword, err := db.Prepare("SELECT password FROM user WHERE pseudo = ?")
	if err != nil {
		db.Close()
		return false, err
	}
	verifPassword, err := selectPassword.Query(pseudo)
	if err != nil {
		db.Close()
		return false, err
	}
	for verifPassword.Next() {
		verifPassword.Scan(&eachPseudo.Password)
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.Password), []byte(password))
		if err != nil {
			db.Close()
			return false, err
		} else {
			verifPassword.Close()
			return true, nil
		}
	}
	db.Close()
	return false, nil
}
