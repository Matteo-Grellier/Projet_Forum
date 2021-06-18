package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	BDD "../BDD"
)

type catName struct {
	Name string
}

func Home(w http.ResponseWriter, req *http.Request) {
	//Tableau regroupant les URLS du site
	arr := []string{"/", "/connexion", "/likedPosts", "/postsActivity", "/topic", "/inscription", "/test"}
	// Ajout de l'url pour chaque catégorie dans le tableau
	pattern := regexp.MustCompile(`\d+`)
	findString := pattern.FindString(req.URL.Path)
	arr = append(arr, "/oneCategory="+findString)
	URLinString := string(req.URL.Path)
	re, err := regexp.Compile("/oneCategory")

	found := re.MatchString(URLinString)

	if err != nil {
		log.Fatal(err)
	}
	if !found {
		temp := "./templates/home.html"
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'home'")
		wichTemplate(w, req, arr, temp)
	} else {
		temp := "./templates/one_category.html"
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'One_category'")
		wichTemplate(w, req, arr, temp)
	}
}

//Fonction qui permet de vérifier si l'url existe sur notre site
func URLfound(arr []string, test string) bool {
	isInArr := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == test {
			isInArr = true
		}
	}
	return isInArr
}

// Fonction qui permet de charger le template home.html soit one_category.html
func wichTemplate(w http.ResponseWriter, req *http.Request, arr []string, temp string) {
	t, err := template.ParseFiles(temp, "./templates/layouts/sidebar.html", "./templates/layouts/header.html", "./templates/layouts/actus.html")
	fmt.Println("je passe par la")
	test := string(req.URL.Path)
	// Si l'url existe alors continuer
	if URLfound(arr, test) {
		fmt.Println("Page Home ✅")
		var nameElement catName
		db := BDD.OpenDataBase()
		pattern := regexp.MustCompile(`\d+`)
		findString := pattern.FindString(req.URL.Path)
		fmt.Println(findString)
		name, err := db.Query("SELECT name FROM category WHERE ID =" + findString)
		for name.Next() {
			name.Scan(&nameElement.Name)
		}
		if err != nil {
			Color(4, "[BDD_INFO] : 🔻 Error BDD : ")
			log.Fatal(err)
		}
		fmt.Println(nameElement.Name)
		t.Execute(w, nameElement)
		// Si l'url n'existe pas, charger la page 404
	} else {
		t, _ = template.ParseFiles("./templates/layouts/error404.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page404'")
		t.Execute(w, nil)
		return
	}
	if err != nil {
		t, _ = template.ParseFiles("./templates/layouts/error500.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page500'")
		t.Execute(w, nil)
		return
	}

}
