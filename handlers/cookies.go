package handlers

import (
	"fmt"
	"time"
	"net/http"
	guuid "github.com/google/uuid"
	BDD "../BDD"
)

var c http.Cookie

func CreateCookie(w http.ResponseWriter, r *http.Request) string{
	id := guuid.New()

	start := time.Now()
	start2 := start.Add(time.Minute*10)
	c = http.Cookie{
			Name:   "ithinkidroppedacookie",
			Value:  id.String(),
			Expires: start2}
	http.SetCookie(w, &c)

	fmt.Printf("%T", c)

	return id.String()
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c = http.Cookie{
			Name:   "ithinkidroppedacookie",
			MaxAge: -1}
	http.SetCookie(w, &c)
}

func GetDeconnected(w http.ResponseWriter, r *http.Request){
	DeleteCookie(w, r)
	BDD.DeleteUUID(UUID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func VerifyCookie() bool{
	if (c.Value == UUID){
		fmt.Println("OK")
		return true
	}
	return false
}