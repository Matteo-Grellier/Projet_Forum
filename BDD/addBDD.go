package BDD

import (
	"database/sql"
	"log"
)

func AddUser(pseudo string, mail string, password string) {
	db := OpenDataBase()
	createNew, err := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = createNew.Exec(pseudo, mail, password)
	if err != nil {
		log.Fatal(err)
	}
}

// Ajout de l'UUID dans la base de données
func AddUUID(pseudo string, UUID string) {
	db := OpenDataBase()
	var actionBDD *sql.Stmt
	var errAction error

	// Vérification de la présence de l'utilisateur dans la table session
	correctPseudo, _ := VerifyBDD(pseudo, "session")
	if correctPseudo {
		actionBDD, errAction = db.Prepare("UPDATE session SET UUID = ? WHERE user_pseudo = ?")
	} else {
		actionBDD, errAction = db.Prepare("INSERT INTO session (UUID, user_pseudo) VALUES(?, ?)")
	}
	if errAction != nil {
		log.Fatal(errAction)
	}

	//Ajout ou update de l'UUID à l'utilisateur connecté
	_, err := actionBDD.Exec(UUID, pseudo)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func AddTopic(title string, content string, user_pseudo string, categoryID int) {
	db := OpenDataBase()
	createNew, err3 := db.Prepare("INSERT INTO topic (title, content, user_pseudo, category_id) VALUES (?, ?, ?, ?)")
	if err3 != nil {
		// Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err3)
	}
	createNew.Exec(title, content, user_pseudo, categoryID)
	createNew.Exec("commit")
	db.Close()
}
func AddLike(user_pseudo string, post_ID int, statusLike int) {
	db := OpenDataBase()
	var actionBDD *sql.Stmt
	var errAction error
	correctPseudo, _ := VerifyBDD(user_pseudo, "like")
	if correctPseudo {
		actionBDD, errAction = db.Prepare("UPDATE like SET liked = ? WHERE user_pseudo = ? AND post_id = ?")
	} else {
		actionBDD, errAction = db.Prepare("INSERT INTO like (liked, user_pseudo, post_id) VALUES(?, ?, ?)")
	}
	if errAction != nil {
		log.Fatal(errAction)
	}

	//Ajout ou update de l'UUID à l'utilisateur connecté
	_, err := actionBDD.Exec(statusLike, user_pseudo, post_ID)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
