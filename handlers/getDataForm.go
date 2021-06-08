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

func GetRegister(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	pseudo := r.FormValue("Pseudo")
	email := r.FormValue("Email")
	password := r.FormValue("Password")
	confirmPwd := r.FormValue("ConfirmPassword")

	/* var allDataRegister = []string{pseudo, email, password, confirmPwd}
	verifyInput(allDataRegister) */

	fmt.Println(pseudo, password, confirmPwd, email)
	var newPass = hashPassword(password)
	fmt.Println(newPass)
	http.Redirect(w, r, "/connexion", http.StatusSeeOther)
}

func hashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

/* func verifyInput(label []string) {

	for index := 0; index < len(label); index++ {
		if len(label[index]) == 0 {
			fmt.Println("Erreur")

		}
	}
} */