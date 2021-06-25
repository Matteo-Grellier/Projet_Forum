package BDD

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Afficher(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/BDD.html")
	DataUsedOK := DataUsed{
		Users:  SelectUsers(),
		Topics: SelectTopics(),
	}

	if req.Method == "POST" {
		if req.FormValue("delete") == "delete" {
			delete()
		} else if req.FormValue("create") == "create" {
			create()
		}
	}
	t.Execute(w, DataUsedOK)
}

func OpenDataBase() *sql.DB {
	/*Ouverture de la base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if err != nil {
		fmt.Println("Could Not open Database")
	}
	return db
}

func SelectUsers() []User {
	/*Sélection du pseudo et de l'adresse mail de tous les utilisateurs*/
	db := OpenDataBase()
	var eachUser User
	var tabUsers []User
	entries, err := db.Query("SELECT pseudo, mail FROM user")

	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for entries.Next() {
		entries.Scan(&eachUser.Pseudo, &eachUser.Mail)
		tabUsers = append(tabUsers, eachUser)
	}
	return tabUsers
}
func SelectTopics() []Topic {
	db := OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	topics, err := db.Query("SELECT * FROM topic")
	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.ID, &eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo, &eachTopic.Category_ID)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics

}

func create() {
	db := OpenDataBase()
	creation, _ := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES(?, ?, ?)")
	_, err := creation.Exec("TEST", "mail@gmail.com", "password")
	if err != nil {
		log.Fatal(err)
	}
}

func delete() {
	db := OpenDataBase()
	delete, _ := db.Prepare("DELETE FROM ? WHERE ? = ?")
	_, err := delete.Exec("TEST", "mail@gmail.com", "password")
	if err != nil {
		log.Fatal(err)
	}
	// Modifier les ? en fonction de ce qu'on veut supprimer
}
