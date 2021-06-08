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
	verifMdp(password)
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

func verifMdp(mdp string) bool {
	var maj int = 0
	var min int = 0
	var chiffre int = 0

	for index := 0; index < len(mdp); index++ {
		if mdp[index] >= 65 && mdp[index] <= 90 {
			maj++
		}
		if mdp[index] >= 97 && mdp[index] <= 122 {
			min++
		}
		if mdp[index] >= 48 && mdp[index] <= 57 {
			chiffre++
		}
	}
	if len(mdp) >= 7 {
		fmt.Println("CA MARCHE")
		if maj >= 1 {
			fmt.Println("Maj Présente")
		} else {
			fmt.Println("PAS DE MAJ")
			return false
		}
		if min >= 1 {
			fmt.Println("Min présente")
		} else {
			fmt.Println("Pas de min")
			return false
		}
		if chiffre >= 1 {
			fmt.Println("Chiffre présent")
		} else {
			fmt.Println("Pas de chiffre")
			return false
		}
	} else {
		fmt.Println("CA MARCHE PAS DEBILE")
		return false
	}
	return true
}
