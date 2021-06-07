package BDD 

import (
	"text/template"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"net/http"
	"log"
)

type User struct{
	Pseudo string
	Mail string
}
type DataUsed struct {
	Users []User
}
func Afficher(w http.ResponseWriter, req *http.Request){
	t, _ := template.ParseFiles("./templates/index.html")
	DataUsedOK := DataUsed{
		Users : SelectUsers(),
	}
	fmt.Println(DataUsedOK)
	t.Execute(w, DataUsedOK)
}

func OpenDataBase() *sql.DB {
	/*Ouverture de la base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if (err!=nil){
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

	if (err!=nil){
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for entries.Next(){
		entries.Scan(&eachUser.Pseudo, &eachUser.Mail)
		tabUsers = append(tabUsers, eachUser)
	}
	return tabUsers
}

// func update(){
// 	db := OpenDataBase()
// 	update, err := db.Prepare("UPDATE user SET pseudo = ? WHERE pseudo = ?")
// 	update.Exec("testUdpate", "test3")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else if update == nil {
// 		fmt.Println(update)
// 	}
// }

// func create(inputPseudo string, inputMail string, inputPassword string){
// 	db := OpenDataBase()
// 	creation, _ := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES(?, ?, ?)")
// 	creation.Exec(inputPseudo, inputMail, inputPassword)
// }

// func delete(inputMail string){
// 	db := OpenDataBase()
// 	delete, _ := db.Prepare("DELETE FROM user WHERE mail = ?")
// 	delete.Exec(inputMail)
// }
