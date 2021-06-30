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

	// On regarde dans quelle table | colonne on souhaite faire la requête
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
		// On stocke l'élément trouvé et on vérifie s'il correspond à celui donné.
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

// Vérifie si l'utilisateur à déjà liké un post
func VerifyLike(post_id int, user_pseudo string) (bool, int, error) {
	var oneUser string
	var statusLike int
	db := OpenDataBase()
	// On prépare notre requête à la table "like" avec les colonnes qui nous intéressent.
	prepareLike, err := db.Prepare("SELECT user_pseudo, liked FROM like WHERE post_id = ?")
	if err != nil {
		db.Close()
		return false, 0, err
	}
	// On lance notre requête au post correspondant
	allLikes, err := prepareLike.Query(post_id)
	if err != nil {
		db.Close()
		return false, 0, err
	}
	for allLikes.Next() {
		// On vérifie si l'utilisateur du like est celui qu'on donne
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

	// On prépare notre requête à la table "user" pour récupérer le mot de passe
	selectPassword, err := db.Prepare("SELECT password FROM user WHERE pseudo = ?")
	if err != nil {
		db.Close()
		return false, err
	}
	// On va chercher le mot de passe à l'utilisateur demandé
	verifPassword, err := selectPassword.Query(pseudo)
	if err != nil {
		db.Close()
		return false, err
	}
	for verifPassword.Next() {
		// On stocke le mot de passe et on compare les 2 mots de passe hashés
		verifPassword.Scan(&eachPseudo.Password)
		err := bcrypt.CompareHashAndPassword([]byte(eachPseudo.Password), []byte(password))
		if err != nil {
			db.Close()
			return false, nil
		} else {
			verifPassword.Close()
			return true, nil
		}
	}
	db.Close()
	return false, nil
}
