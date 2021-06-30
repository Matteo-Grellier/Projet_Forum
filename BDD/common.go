package BDD

import (
	"database/sql"
	"fmt"
)

// Fonction pour ouvrir la Base De Données
func OpenDataBase() *sql.DB {
	/*Ouverture de la base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if err != nil {
		fmt.Println("Could Not open Database")
	}
	return db
}

// Définition des structures utilisées pour les requêtes à la BDD

type User struct {
	Pseudo   string
	Mail     string
	Password string
}

type Category struct {
	Name string
	Id   string
}

type Topic struct {
	ID            int
	Title         string
	Content       string
	Like          int
	User_pseudo   string
	Category_ID   int
	Category_name string
}

type Post struct {
	ID             int
	Content        string
	User_pseudo    string
	Topic_ID       string
	Comments       []Comment
	NumberComments int
	NumberLikes    int
	NumberDislikes int
	UserLiked      bool
	UserDisliked   bool
}

type Comment struct {
	ID          int
	User_pseudo string
	Content     string
	Post_ID     int
}

type Likes struct {
	Status      int
	User_Pseudo string
	ID          int
}
