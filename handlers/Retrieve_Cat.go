package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	Name string
}

type Data struct {
	Categories []Category
}

func RetrieveCat(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/all_categories.html")

	/*fonction base de données*/
	db, err := sql.Open("sqlite3", "./BDD/BDDForum1.db")
	if err != nil {
		fmt.Println("Could Not open Database")
	}

	dataOk := Data{
		Categories: bdd(db),
	}
	fmt.Println(dataOk)
	t.Execute(w, dataOk)
}

func bdd(db *sql.DB) []Category {

	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name FROM category")

	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
		// return
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name)
		tabCategories = append(tabCategories, eachCategory)
		fmt.Println(tabCategories)
	}

	return tabCategories
}
