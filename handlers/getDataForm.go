package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	BDD "../BDD"
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
	if isValidEmail(email) && verifyBDD(email, "mail") && verifyBDD(pseudo, "pseudo"){
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	}
}

func hashPassword(password string) string {
	var passByte = []byte(password)

	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
func isValidEmail(email string) bool{
	// Vérifie si l'email entré est valide
	var re = regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)
	if re.MatchString(email) {
		return true
	}
	fmt.Println("Email non correct.")
	return false
}

func verifyBDD(element string, column string) bool {
	// Vérifie si l'élément est déjà dans la base de donnée ou pas.
	db := BDD.OpenDataBase()
	var oneElement string
	var tabElements []string
	allElements, err := db.Query("SELECT " + column + " FROM user")
	if (err!=nil){
		fmt.Println("Could not query database")
		log.Fatal(err)
	}
	for allElements.Next(){
		allElements.Scan(&oneElement)
		tabElements = append(tabElements, oneElement)
	}
	for _,eachElement := range tabElements{
		if (eachElement == element) {
			fmt.Println(column + " déjà dans la base de données.")
			// créer une erreur qui lance la template !
			return false
		}
	}

	return true
}

/*func verifyInput(label []string) {

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
