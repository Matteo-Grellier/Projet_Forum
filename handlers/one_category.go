package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	BDD "../BDD"
)

type DataUsed struct {
	Topics       []Topic
	Category     string
	CategoryID   int
	ErrorMessage string
}

type Topic struct {
	Title       string
	Content     string
	User_pseudo string
}

func One_Category(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("cat"))
	t, err := template.ParseFiles("templates/one_category.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")

	var DataUsedOK DataUsed

	DataUsedOK.ErrorMessage = ""

	if r.Method == "POST" {
		DataUsedOK.ErrorMessage = GetTopic(w, r).Error
	}

	DataUsedOK.Category = DisplayCategory(w, r, categoryID)
	DataUsedOK.Topics = DisplayTopics(categoryID)
	DataUsedOK.CategoryID = categoryID

	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'one_category'")
	t.Execute(w, DataUsedOK)
}

func DisplayTopics(idCat int) []Topic {
	db := BDD.OpenDataBase()
	var eachTopic Topic
	var tabTopics []Topic
	searchTopics, err := db.Prepare("SELECT title, content, user_pseudo FROM topic WHERE category_id = ?")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}

	topics, err := searchTopics.Query(idCat)
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for topics.Next() {
		topics.Scan(&eachTopic.Title, &eachTopic.Content, &eachTopic.User_pseudo)
		tabTopics = append(tabTopics, eachTopic)
	}
	return tabTopics
}

func DisplayCategory(w http.ResponseWriter, r *http.Request, idCat int) string {
	var nameElement string
	db := BDD.OpenDataBase()
	searchName, err := db.Prepare("SELECT name FROM category WHERE ID = ?")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	categoryFound, err := searchName.Query(idCat)
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for categoryFound.Next() {
		categoryFound.Scan(&nameElement)
	}
	if nameElement == "" {
		// Ajouter la fonction d'erreur si l'ID n'est pas valide
		http.Redirect(w, r, "/oneCategory?cat=0", http.StatusSeeOther)
	}

	return nameElement
}

func GetTopic(w http.ResponseWriter, r *http.Request) Errors {
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("cat"))
	err2 := r.ParseForm()
	if err2 != nil {
		Color(4, "[PARSE_FORM_INFO] : 🔻 Error function 'GetTopic' : ")
		log.Fatal(err2)
	}

	titre := r.FormValue("titre")
	post := r.FormValue("post")
	//TEST BRUT
	user := "L1"
	categId := categoryID

	var data = []string{titre, post}

	var ErrorsPost Errors

	if verifyInput(data) {
		db := BDD.OpenDataBase()
		createNew, err3 := db.Prepare("INSERT INTO topic (title, content, user_pseudo, category_id) VALUES (?, ?, ?, ?)")
		if err3 != nil {
			Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
			log.Fatal(err3)
		}
		createNew.Exec(titre, post, user, categId)
	} else {
		ErrorsPost.Error = ErrorMessage
		ErrorMessage = ""
	}
	return ErrorsPost
}
