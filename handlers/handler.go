package handlers

import (
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
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'inscription'")
	t.Execute(w, nil)
}

//Exécution de la page Categories
func CategoriesPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/categories.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")

	userConnected := VerifyUserConnected(w, req)
	if err != nil {
		Error500(w, req, err)
	}
	type TabCategories struct {
		Categories []BDD.Category
		User       UserConnectedStruct
	}

	dataOk := TabCategories{
		Categories: BDD.DisplayCategories(),
		User:       userConnected,
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Catégories'")
	t.Execute(w, dataOk)
}

//Exécution de la page oneCategory
func OneCategoryPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/oneCategory.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")

	if err != nil {
		Error500(w, r, err)
		// Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		// log.Fatalf("%s", err)
		return
	}

	if !Error404(w, r) {
		return
	}
	// Récupération de l'ID de la catégorie
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("cat"))

	DataPageCategoryOK := DataPageCategory{
		Category:      BDD.DisplayCategory(categoryID),
		Topics:        BDD.DisplayTopics(categoryID),
		CategoryID:    categoryID,
		UserConnected: VerifyUserConnected(w, r),
	}
	if DataPageCategoryOK.Category == "nil" {
		NoItemsError(w)
		return
	}

	if r.Method == "POST" {

		// Vérification du cookie du navigateur
		if VerifyUserConnected(w, r).Connected {
			pseudo := VerifyUserConnected(w, r).PseudoConnected
			var topicID int
			titre := r.FormValue("titre")
			content := r.FormValue("post")

			// Ajout du topic OU affichage de l'erreur
			DataPageCategoryOK.Error, topicID = AddTopic(titre, content, categoryID, pseudo)
			DataPageCategoryOK.Topics = BDD.DisplayTopics(categoryID)
			if DataPageCategoryOK.Error != "" {
				t.Execute(w, DataPageCategoryOK)
				return
			} else {
				http.Redirect(w, r, "/topic?top="+strconv.Itoa(topicID), http.StatusSeeOther)
			}
		} else {
			DataPageCategoryOK.Error = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un topic."
		}
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'one_category'")
	t.Execute(w, DataPageCategoryOK)
}
func OneTopicPage(w http.ResponseWriter, r *http.Request) {

	TopicID, _ := strconv.Atoi(r.URL.Query().Get("top"))
	t, err := template.ParseFiles("templates/topic.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/boxPost.html", "./templates/layouts/boxComm.html")
	var DataPageTopicOK TopicDataUsed
	DataPageTopicOK.ErrorMessage = ""

	DataPageTopicOK = TopicDataUsed{
		ErrorMessage:  "",
		Topic:         BDD.DisplayOneTopic(TopicID),
		Posts:         BDD.DisplayPosts(TopicID),
		UserConnected: VerifyUserConnected(w, r),
	}
	if r.Method == "POST" {
		if r.FormValue("Post") != "" {
			if VerifyUserConnected(w, r).Connected {
				userPseudo := VerifyUserConnected(w, r).PseudoConnected
				postContent := r.FormValue("Post")
				BDD.AddPost(userPseudo, postContent, TopicID)
			} else {
				DataPageTopicOK.ErrorMessage = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un post."
			}
		} else if r.FormValue("Comment") != "" {
			if VerifyUserConnected(w, r).Connected {
				userPseudo := VerifyUserConnected(w, r).PseudoConnected
				comment := r.FormValue("Comment")
				postID, _ := strconv.Atoi(r.FormValue("postID"))
				BDD.AddComment(comment, userPseudo, postID)
			} else {
				DataPageTopicOK.ErrorMessage = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un commentaire."
			}
		}
		DataPageTopicOK.Posts = BDD.DisplayPosts(TopicID)
	}
	if err != nil {
		Error500(w, r, err)
		// Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		// log.Fatalf("%s", err)
		return
	}

	if !Error404(w, r) {
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")
	t.Execute(w, DataPageTopicOK)
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
