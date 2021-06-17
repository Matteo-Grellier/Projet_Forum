package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	BDD "../BDD"
	guuid "github.com/google/uuid"
)

var UUIDConnected string

func CreateCookie(w http.ResponseWriter, r *http.Request) string {
	id := guuid.New()

	start := time.Now()
	start2 := start.Add(time.Second * 10)
	c := http.Cookie{
		Name:    "CookieSession",
		Value:   id.String(),
		Expires: start2}
	http.SetCookie(w, &c)
	UUIDConnected = c.Value
	fmt.Printf("%T", c)

	return id.String()
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "CookieSession",
		MaxAge: -1}
	http.SetCookie(w, &c)
}

func GetDeconnected(w http.ResponseWriter, r *http.Request) {
	BDD.DeleteUUID(UUIDConnected)
	UUIDConnected = ""
	DeleteCookie(w, r)
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	CookieValue := ReadCookie(w, r, "CookieSession")
	var oneUUID string
	db := BDD.OpenDataBase()
	checkUUID, err := db.Prepare("SELECT UUID FROM session WHERE user_pseudo = ?")
	if err != nil {
		log.Fatal(err)
	}
	UUIDPseudo, err := checkUUID.Query("Roberto04")
	if err != nil {
		log.Fatal(err)
	}

	for UUIDPseudo.Next() {
		UUIDPseudo.Scan(&oneUUID)
		if oneUUID == UUID {
			fmt.Println(oneUUID)
		}
	}
	checkUUID.Close()

	if CookieValue == UUID {
		fmt.Println("OK")
		fmt.Println("Valeur du cookie", CookieValue)
		fmt.Println("Valeur de l'UUID", UUID)

		return true
	}
	return false
}

func ReadCookie(w http.ResponseWriter, r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		GetDeconnected(w, r)
		return "undefined"
	}
	println("\033[0;32m", "[cookies] : here are the chocolat chips in your cookies :", c.Value)
	return c.Value
}
