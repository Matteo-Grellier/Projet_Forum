package handlers 

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

type Data struct {
	Users []User
}
func Afficher(w http.ResponseWriter, req *http.Request){
	t, _ := template.ParseFiles("./template/index.html")

	/*fonction base de données*/
	db, err := sql.Open("sqlite3", "../BDDForum1.db")
	if (err!=nil){
		fmt.Println("Could Not open Database")
	}

	if req.Method == "GET"{
		if req.FormValue("update")=="update" {
			update(db)
		} else if req.FormValue("create") == "create"{
			create(db)
		}  else if req.FormValue("delete") == "delete"{
			delete(db)
		}
	}

	dataOk := Data{
		Users : bdd(db),
	}
	t.Execute(w, dataOk)
}

func update(db *sql.DB){
	update, err := db.Prepare("UPDATE user SET pseudo = ? WHERE pseudo = ?")
	update.Exec("liv44", "liv49")
	if err != nil {
		fmt.Println(err)
	} else if update == nil {
		fmt.Println(update)
	}
}

func create(db *sql.DB){
	creation, _ := db.Prepare("INSERT INTO user (pseudo, mail, password) VALUES(?, ?, ?)")
	creation.Exec("Mat", "mat@gmail.com", "qwerty")
}

func delete(db *sql.DB){
	delete, _ := db.Prepare("DELETE FROM user WHERE pseudo = ?")
	delete.Exec("Mat")

}

func bdd(db *sql.DB) []User {
	
	var eachUser User
	var tabUsers []User
	entries, err := db.Query("SELECT * FROM post")

	if (err!=nil){
		fmt.Println("Could not query database")
		log.Fatal(err)
		// return
	}
	for entries.Next(){
		entries.Scan(&eachUser.Pseudo, &eachUser.Mail)
		tabUsers = append(tabUsers, eachUser)
	}

	return tabUsers
}