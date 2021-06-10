package handlers

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	pseudo := r.FormValue("Pseudo")
	password := r.FormValue("Password")

	/* var allDataLogin = []string{pseudo, password}
	verifyInput(allDataLogin) */

	fmt.Println(pseudo, password)
	var newPass = hashPassword(password)
	fmt.Println(newPass)
	http.Redirect(w, r, "/all_categories", http.StatusSeeOther)
}

func hashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

