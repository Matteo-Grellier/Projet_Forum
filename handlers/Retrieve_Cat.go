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
	t, _ := template.ParseFiles("./templates/one_category.html")

	/*fonction base de données*/
	db := BDD.OpenDataBase()

	dataOk := Data{
		Categories: bdd(db),
	}
	bar := req.FormValue("foo")
	if bar != "" {
		finalURL := "oneCategory=" + bar
		http.Redirect(w, req, finalURL, http.StatusSeeOther)
	}
	t.Execute(w, dataOk)
}

func bdd(db *sql.DB) []Category {

	var eachCategory Category
	var tabCategories []Category
	entries, err := db.Query("SELECT name,ID FROM category")

	if err != nil {
		fmt.Println("Could not query database")
		log.Fatal(err)
		// return
	}
	for entries.Next() {
		entries.Scan(&eachCategory.Name, &eachCategory.Id)
		tabCategories = append(tabCategories, eachCategory)
	}
	return tabCategories
}
