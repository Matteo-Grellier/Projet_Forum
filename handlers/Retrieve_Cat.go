package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	Name string
	Id   string
}

type Data struct {
	Categories []Category
}

func RetrieveCat(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/all_categories.html")

	/*fonction base de données*/
	db := BDD.OpenDataBase()

	dataOk := Data{
		Categories: bdd(db),
	}

	if req.Method == "POST" {
		fmt.Println("ON clique sur une catégorie")
		categorie := req.FormValue("foo")
		fmt.Println(categorie)
		One_Category(w, req)
	} else {
		t.Execute(w, dataOk)
	}



func bdd(db *sql.DB) []Category {

	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name,ID FROM category")

	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
		// return
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name, &eachCategory.Id)
		tabCategories = append(tabCategories, eachCategory)
	}
	return tabCategories
}
