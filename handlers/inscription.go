package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"regexp"
	BDD "../BDD"
)

var ErrorMessage string
type Errors struct{
	Error string
	Pseudo string
	Mail string
}

func InscriptionPage(w http.ResponseWriter, r *http.Request) {
	// Déclaration des fichiers à parser
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	fmt.Printf("%T\n", t)

	fmt.Println("Page Inscription ✅")
	t.Execute(w, nil)
}

func GetRegister(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/inscription.html", "./templates/layouts/sidebar.html", "./templates/layouts/header.html")
	if err != nil {
		log.Fatalf("Template execution: %s", err)
		return
	}
	err2 := r.ParseForm()
	if err2 != nil {
		log.Fatal(err2)
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
	if isValidEmail(email) && verifyBDD(email, "mail") && verifyBDD(pseudo, "pseudo") && verifMdp(password){
		http.Redirect(w, r, "/connexion", http.StatusSeeOther)
	} else {
		fmt.Println(ErrorMessage)
		ErrorsInscriptions := Errors{
			Error : ErrorMessage,
			Pseudo : pseudo,
			Mail : email,
		}
		ErrorMessage = ""	
		t.Execute(w, ErrorsInscriptions)
	}
}

func isValidEmail(email string) bool{
	// Vérifie si l'email entré est valide
	var re = regexp.MustCompile(`(?mi)[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}`)
	if re.MatchString(email) {
		return true
	}
	ErrorMessage = "Email non valide."
	// fmt.Println("Email non correct.")
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
			ErrorMessage = column + " déjà dans la base de données."
			// fmt.Println(column + " déjà dans la base de données.")
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
		ErrorMessage = "Mot de passe non valide."
		return false
	}
	return true
}

/* func getPassword(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var e string = r.PostFormValue("Password")
	fmt.Println(string("\033[1;37m\033[0m"), e)
}
*/
