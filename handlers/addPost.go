package handlers

import (
	"log"
	"net/http"
	"text/template"

	BDD "../BDD"
)

func Post(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("")
	if err != nil {
		Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
		log.Fatalf("%s", err)
		return
	}

	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'addTopic'")
	t.Execute(w, nil)
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		Color(4, "[PARSE_FORM_INFO] : 🔻 Error function 'GetValue' : ")
		log.Fatal(err)
	}

	pseudo := "Eloulou2001"
	post := r.FormValue("Post")
	// L'id et le pseudo a supp car cela deviendra automatique par la suite.
	id := 128

	addPost(pseudo, post, id)
}

func addPost(pseudo string, post string, id int) {
	db := BDD.OpenDataBase()
	add, _ := db.Prepare("INSERT INTO post (user_pseudo, content, topic_id) VALUES (?, ?, ?)")
	_, err := add.Exec(pseudo, post, id)
	if err != nil {
		Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
		log.Fatalf("%s", err)
	}
}
