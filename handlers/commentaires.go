package handlers

import (
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

type DataComment struct {
	Comments     []Comment
	ErrorMessage string
}

type Comment struct {
	Content     string
	User_pseudo string
}

func Commentaires(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/commentaires.html")

	var DataCommentOK DataComment

	DataCommentOK.ErrorMessage = ""

	if r.Method == "POST" {
		Comment := r.FormValue("comment")
		pseudo := "Roberto04"
		addComment(Comment, pseudo)
	}

	DataCommentOK.Comments = DisplayComments()

	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'posts_activity'")
	t.Execute(w, DataCommentOK)
}

func DisplayComments() []Comment {
	db := BDD.OpenDataBase()
	var eachComment Comment
	var tabComments []Comment
	comments, err := db.Query("SELECT content, user_pseudo FROM comment")
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatal(err)
	}
	for comments.Next() {
		comments.Scan(&eachComment.Content, &eachComment.User_pseudo)
		tabComments = append(tabComments, eachComment)
	}
	return tabComments
}

func addComment() {

}
