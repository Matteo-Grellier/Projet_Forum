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
