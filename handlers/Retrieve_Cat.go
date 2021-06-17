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
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "[BDD_INFO] : 🔹", dataOk)
	t.Execute(w, dataOk)
}

func bdd(db *sql.DB) []Category {

	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name FROM category")

	if err != nil {
		colorYellow := "\033[33m"
		log.Fatalf(string(colorYellow), "[BDD_INFO] : 🔻 Template execution: %s", err)
		// return
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name)
		tabCategories = append(tabCategories, eachCategory)
		fmt.Println(tabCategories)
	}

	return tabCategories
}
