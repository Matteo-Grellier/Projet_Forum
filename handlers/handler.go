package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	BDD "../BDD"
	_ "github.com/mattn/go-sqlite3"
)

type Errors struct {
	Error  string
	Pseudo string
	Mail   string
}

var ErrorMessage string

// Exécution de la page Home
func Home(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/bouton_all_categories.html", "./templates/layouts/actus.html")

	userConnected := VerifyUserConnected(w, req)
	if !Error404(w, req) {
		return
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page500'")
		t.Execute(w, nil)
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'home'")
	t.Execute(w, userConnected)
}

// Exécution de la page Connexion
func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/connection.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Error500(w, r, err)
		// Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		// log.Fatalf("%s", err)
		return
	}
	if !Error404(w, r) {
		return
	}

	if r.Method == "POST" {
		pseudo := r.FormValue("Pseudo")
		password := r.FormValue("Password")
		statusConnexion := GetLogin(w, r, pseudo, password)
		if statusConnexion.Error == "" {
			CreateCookie(w, r, pseudo)
			// CreateUUID(pseudo, UUID)
			Color(1, "[CONNEXION] : 🟢 Vous êtes connecté ")
			http.Redirect(w, r, "/categories", http.StatusSeeOther)
		} else {
			Color(4, "[CONNEXION] : 🔻 Vous n'êtes pas connecté ")
			t.Execute(w, statusConnexion)
		}
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'connexion'")
	t.Execute(w, nil)
}

//Exécution de la page Inscription
func InscriptionPage(w http.ResponseWriter, r *http.Request) {

	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Error500(w, r, err)
		// Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		// log.Fatalf("%s", err)
		return
	}
	if !Error404(w, r) {
		return
	}
	if r.Method == "POST" {
		pseudo := r.FormValue("Pseudo")
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		confirmPwd := r.FormValue("ConfirmPassword")
		statusRegister := GetRegister(pseudo, email, password, confirmPwd)
		if statusRegister.Error == "" {
			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
		} else {
			t.Execute(w, statusRegister)
		}
		fmt.Println("ON S'ENREGISTRE")
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'inscription'")
	t.Execute(w, nil)
}

//Exécution de la page Categories
func CategoriesPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/categories.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Error500(w, req, err)
	}

	type Data struct {
		Categories []BDD.Category
	}

	dataOk := Data{
		Categories: BDD.DisplayCategories(),
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Catégories'")
	t.Execute(w, dataOk)
}

//Exécution de la page oneCategory
func OneCategoryPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("cat"))
	t, err := template.ParseFiles("templates/oneCategory.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")

	var DataUsedOK DataUsed

	DataUsedOK.ErrorMessage = ""

	if r.Method == "POST" {

		DataUsedOK.ErrorMessage = GetTopic(w, r).Error
	}
	DataUsedOK = DataUsed{
		ErrorMessage: "",
		Category:     DisplayCategory(w, r, categoryID),
		Topics:       DisplayTopics(categoryID),
		CategoryID:   categoryID,
	}

	if err != nil {
		Error500(w, r, err)
		// Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		// log.Fatalf("%s", err)
		return
	}
	if DataUsedOK.Category == "nil" {
		NoItemsError(w)
		return
	}
	if !Error404(w, r) {
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'one_category'")
	t.Execute(w, DataUsedOK)
}

//Exécution de la page Topic
func TopicPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/topic.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")
	t.Execute(w, nil)
}

func LikedPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/LikedPage.html", "templates/layouts/sidebar.html")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")
	t.Execute(w, nil)
}
