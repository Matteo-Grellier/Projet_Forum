package BDD

import (
	"database/sql"
)

// Fonction permettant d'ajouter l'utilisateur dans la base de donnée
func AddUser(pseudo string, mail string, password string) error {
	db := OpenDataBase()
	// On prépare la table "user" en indiquant dans quelles colonnes nous stockons les données reçues
	createNew, err := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES (?, ?, ?)")

	if err != nil {
		db.Close()
		return err
	}
	// On ajoute nos données envoyées par le formulaire
	_, err = createNew.Exec(pseudo, mail, password)
	if err != nil {
		db.Close()
		return err
	}
	return nil
}

// Ajout de l'UUID dans la base de données
func AddUUID(pseudo string, UUID string) error {
	db := OpenDataBase()
	var actionBDD *sql.Stmt
	var errAction error

	// Vérification de la présence de l'utilisateur dans la table session
	correctPseudo, _, err := VerifyBDD(pseudo, "session")

	if err != nil {
		db.Close()
		return err
	}
	// Si l'utilisateur est déjà dans la table, on update la donnée, sinon on ajoute un élément
	if correctPseudo {
		actionBDD, errAction = db.Prepare("UPDATE session SET UUID = ? WHERE user_pseudo = ?")
	} else {
		actionBDD, errAction = db.Prepare("INSERT INTO session (UUID, user_pseudo) VALUES(?, ?)")
	}
	if errAction != nil {
		db.Close()
		return errAction
	}

	//Ajout ou update de l'UUID à l'utilisateur connecté
	_, err = actionBDD.Exec(UUID, pseudo)
	if err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}

// Fonction permettant d'ajouter un topic dans la base de donnée
func AddTopic(title string, content string, user_pseudo string, categoryID int) error {
	db := OpenDataBase()
	// On prépare la table "topic" en indiquant dans quelles colonnes nous stockons les données reçues
	createNew, err := db.Prepare("INSERT INTO topic (title, content, user_pseudo, category_id) VALUES (?, ?, ?, ?)")

	if err != nil {
		db.Close()
		return err
	}
	// On ajoute nos données envoyées par le formulaire
	createNew.Exec(title, content, user_pseudo, categoryID)
	createNew.Exec("commit")
	db.Close()
	return nil
}

// Fonction permettant d'ajouter un like dans la base de donnée
func AddLike(user_pseudo string, post_ID int, Liked int) error {
	db := OpenDataBase()
	var actionBDD *sql.Stmt
	var errAction error

	// On vérifie si l'utilisateur a déjà liké/disliké le post, on stocke le status du like
	correctPseudo, statusLike, err := VerifyLike(post_ID, user_pseudo)
	if err != nil {
		db.Close()
		return err
	}

	// On prépare la table "like" en indiquant dans quelles colonnes nous stockons les données reçues
	if correctPseudo {
		actionBDD, errAction = db.Prepare("UPDATE like SET liked = ? WHERE user_pseudo = ? AND post_id = ?")
		// On ajoute le status du like en fonction du status du like déjà en place
		if statusLike == 1 && Liked != -1 || statusLike == -1 && Liked != 1 {
			_, err := actionBDD.Exec(0, user_pseudo, post_ID)
			if err != nil {
				db.Close()
				return err
			}
			db.Close()
			return nil
		}
	} else {
		actionBDD, errAction = db.Prepare("INSERT INTO like (liked, user_pseudo, post_id) VALUES(?, ?, ?)")
	}
	if errAction != nil {
		db.Close()
		return errAction
	}

	// On ajoute nos données
	_, err = actionBDD.Exec(Liked, user_pseudo, post_ID)
	if err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}

// Fonction permettant d'ajouter un post dans la base de donnée
func AddPost(pseudo string, post string, id int) error {
	db := OpenDataBase()
	// On prépare la table "post" en indiquant dans quelles listes nous stockons les données reçues
	add, err := db.Prepare("INSERT INTO post (user_pseudo, content, topic_id) VALUES (?, ?, ?)")
	if err != nil {
		db.Close()
		return err
	}
	// On ajoute nos données envoyées par le formulaire
	_, err = add.Exec(pseudo, post, id)
	if err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}

// Fonction permettant d'ajouter un commentaire dans la base de donnée
func AddComment(comment string, user string, postId int) error {
	db := OpenDataBase()
	// On prépare la table "Comment" en indiquant dans quelles listes nous stockons les données reçues
	createNew, err := db.Prepare("INSERT INTO Comment (content, user_pseudo, post_id) VALUES (?, ?, ?)")
	if err != nil {
		db.Close()
		return err
	}
	// On ajoute nos données envoyées par le formulaire
	_, err = createNew.Exec(comment, user, postId)

	if err != nil {
		db.Close()
		return err
	}
	db.Close()
	return nil
}
