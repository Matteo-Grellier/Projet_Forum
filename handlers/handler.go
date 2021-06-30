package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	BDD "../BDD"
	_ "github.com/mattn/go-sqlite3"
)

// Exécution de la page Home
func Home(w http.ResponseWriter, req *http.Request) {
	// La fonction Erreur404 est lancée à la racine du site.
	if !Error404(w, req) {
		return
	}

	// On va chercher les fichiers HTML pour l'affichage de la page
	t, err := template.ParseFiles("./templates/home.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/bouton_all_categories.html", "./templates/layouts/actus.html")

	// On vérifie si l'utilisateur est connecté
	userConnected := VerifyUserConnected(w, req)

	// On affiche les 3 derniers posts publiés
	Posts, BDDerr := BDD.DisplayPostsActus()
	if BDDerr != nil {
		Error500(w, req, BDDerr)
		return
	}

	// On envoie les données utiles
	DataPageHomeOK := DataPageHome{
		UserConnected: userConnected,
		Posts:         Posts,
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page500'")
		t.Execute(w, nil)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'home'")
	t.Execute(w, DataPageHomeOK)
}

// Exécution de la page Connexion
func ConnexionPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/connection.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Error500(w, r, err)
		return
	}

	// SI une requête est envoyée
	if r.Method == "POST" {
		pseudo := r.FormValue("Pseudo")
		password := r.FormValue("Password")
		// On lance la fonction pour connecter un utilisateur avec les données du formulaire
		statusConnexion, BDDerror := GetLogin(w, r, pseudo, password)
		if BDDerror != nil {
			Error500(w, r, BDDerror)
		}
		if statusConnexion.Error == "" {
			// On créé un cookie
			BDDerror = CreateCookie(w, r, pseudo)
			if BDDerror != nil {
				Error500(w, r, BDDerror)
			}
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
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		Error500(w, r, err)
		return
	}

	var statusRegister Errors

	// SI une requête est envoyée
	if r.Method == "POST" {
		pseudo := r.FormValue("Pseudo")
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		confirmPwd := r.FormValue("ConfirmPassword")

		// On lance la fonction permettant d'inscrire un utilisateur avec les données du formulaire
		statusRegister, BDDerror = GetRegister(pseudo, email, password, confirmPwd)
		if BDDerror != nil {
			Error500(w, r, BDDerror)
			return
		}
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

	// On vérifie si l'utilisateur est connecté
	userConnected := VerifyUserConnected(w, req)
	if err != nil {
		Error500(w, req, err)
	}

	allCategories, BDDerror := BDD.DisplayCategories()
	if BDDerror != nil {
		Error500(w, req, BDDerror)
	}

	// On ajoute les catégories du site dans la structure de données envoyées à la page
	DataPageCategoriesOK := DataPageCategories{
		Categories: allCategories,
		User:       userConnected,
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Catégories'")
	t.Execute(w, DataPageCategoriesOK)
}

//Exécution de la page oneCategory
func OneCategoryPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/oneCategory.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html")

	if err != nil {
		Error500(w, r, err)
		return
	}

	// Récupération de l'ID de la catégorie via l'URL
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("cat"))

	// On va chercher le nom de la catégorie
	oneCategory, BDDerror := BDD.DisplayCategory(categoryID)
	if BDDerror != nil {
		Error500(w, r, BDDerror)
		return
	}
	// On va chercher les topics liés à cette catégorie
	allTopics, BDDerror := BDD.DisplayTopics(categoryID)
	if BDDerror != nil {
		Error500(w, r, BDDerror)
		return
	}
	// On ajoute les données envoyées à la page
	DataPageCategoryOK := DataPageCategory{
		Category:      oneCategory,
		Topics:        allTopics,
		CategoryID:    categoryID,
		UserConnected: VerifyUserConnected(w, r),
	}
	if DataPageCategoryOK.Category == "nil" {
		NoItemsError(w, r)
		return
	}

	// Si une requête est envoyée
	if r.Method == "POST" {

		// On vérifie que l'utilisateur est bien connecté
		if DataPageCategoryOK.UserConnected.Connected {
			pseudo := DataPageCategoryOK.UserConnected.PseudoConnected
			var topicID int
			titre := r.FormValue("titre")
			content := r.FormValue("post")

			// On vérifie que les champs ont bien été remplis
			var data = []string{titre, content}
			if verifyInput(data) {
				// On ajoute le topic dans la BDD
				BDDerror = BDD.AddTopic(titre, content, pseudo, categoryID)
				if BDDerror != nil {
					Error500(w, r, BDDerror)
					return
				}
			} else {
				DataPageCategoryOK.Error = "Il manque un item."
				t.Execute(w, DataPageCategoryOK)
				return
			}

			// On va chercher l'ID du topic que l'on vient de créer
			topicID, BDDerror := BDD.DisplayLastTopic()
			if BDDerror != nil {
				Error500(w, r, BDDerror)
				return
			}
			// On redirige vers la nouvelle page de topic créé
			http.Redirect(w, r, "/topic?top="+strconv.Itoa(topicID), http.StatusSeeOther)
		} else {
			DataPageCategoryOK.Error = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un topic."
		}
	}
	if BDDerror != nil {
		Error500(w, r, BDDerror)
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'one_category'")
	t.Execute(w, DataPageCategoryOK)
}

//Exécution d'une page topic
func OneTopicPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/topic.html", "templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/boxPost.html", "./templates/layouts/boxComm.html")
	if err != nil {
		Error500(w, r, err)
		return
	}
	// On va chercher l'ID du topic via l'URL
	TopicID, _ := strconv.Atoi(r.URL.Query().Get("top"))

	// On va chercher les détails du topic demandé
	oneTopic, BDDerror := BDD.DisplayOneTopic(TopicID)
	if BDDerror != nil {
		Error500(w, r, BDDerror)
	}
	DataPageTopicOK := TopicDataUsed{
		ErrorMessage:  "",
		Topic:         oneTopic,
		UserConnected: VerifyUserConnected(w, r),
	}

	// On Déclare une variable qui stocke le pseudo de l'utilisateur connecté
	userPseudo := DataPageTopicOK.UserConnected.PseudoConnected

	// On va chercher tous les posts du topic
	allPosts, BDDerror := BDD.DisplayPosts(TopicID, userPseudo)
	if BDDerror != nil {
		Error500(w, r, BDDerror)
	}
	DataPageTopicOK.Posts = allPosts

	if DataPageTopicOK.Topic.Category_name == "nil" {
		NoItemsError(w, r)
		return
	}
	// Si un requête est envoyée
	if r.Method == "POST" {
		// On vérifie quel formulaire a été envoyé
		if r.FormValue("Post") != "" {
			// On vérifie que l'utilisateur soit connecté avant de pouvoir ajouter un post
			if DataPageTopicOK.UserConnected.Connected {
				postContent := r.FormValue("Post")
				BDDerror = BDD.AddPost(userPseudo, postContent, TopicID)
				if BDDerror != nil {
					Error500(w, r, BDDerror)
					return
				}
			} else {
				DataPageTopicOK.ErrorMessage = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un post."
			}

		} else if r.FormValue("Comment") != "" {
			// On vérifie que l'utilisateur soit connecté avant de pouvoir ajouter un commentaire
			if DataPageTopicOK.UserConnected.Connected {
				comment := r.FormValue("Comment")
				postID, _ := strconv.Atoi(r.FormValue("postID"))
				BDDerror = BDD.AddComment(comment, userPseudo, postID)
				if BDDerror != nil {
					Error500(w, r, BDDerror)
					return
				}
			} else {
				DataPageTopicOK.ErrorMessage = "Vous n'êtes pas connectés. Vous devez vous connecter pour ajouter un commentaire."
			}

		} else if r.FormValue("like") == "0" {
			if DataPageTopicOK.UserConnected.Connected {
				// On vérifie que l'utilisateur soit connecté avant de pouvoir liker/disliker un post
				post_id, _ := strconv.Atoi(r.FormValue("post_id"))
				likeOrDislike, _ := strconv.Atoi(r.FormValue("status"))
				BDDerror = BDD.AddLike(userPseudo, post_id, likeOrDislike)
				if BDDerror != nil {
					Error500(w, r, BDDerror)
					return
				}
			} else {
				DataPageTopicOK.ErrorMessage = "Vous n'êtes pas connectés. Vous devez vous connecter pour liker un post."
			}
		}

		// On recharge les posts, commentaires, likes affichés sur le site si certains ont été ajoutés
		DataPageTopicOK.Posts, BDDerror = BDD.DisplayPosts(TopicID, userPseudo)
		if BDDerror != nil {
			Error500(w, r, BDDerror)
			return
		}
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'topic'")
	t.Execute(w, DataPageTopicOK)
}

func LikesPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/likes.html", "templates/layouts/header.html", "templates/layouts/sidebar.html")
	if err != nil {
		Error500(w, r, err)
		return
	}
	// On regarde si l'utilisateur est connecté
	user_connected := VerifyUserConnected(w, r)
	if !user_connected.Connected {
		Error500(w, r, err)
		return
	}
	// On va chercher les posts likés
	posts, BDDerror := BDD.DisplayLikedPosts(user_connected.PseudoConnected)
	if BDDerror != nil {
		Error500(w, r, BDDerror)
		return
	}
	// On retourne les données utiles pour la page
	DataPageLikesOK := DataPageLikes{
		UserConnected: user_connected,
		Posts:         posts,
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'likes'")
	t.Execute(w, DataPageLikesOK)
}
