package controller

import (
	"golantah/backend"
	templates "golantah/temp"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

var sexe string
var img string

// DefaultHandler est la fonction qui redirige vers la page 404 en cas de route inconnue
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusNotFound)
	err := templates.Temp.ExecuteTemplate(w, "erreur", nil)
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "index", nil)
}

func DisplayChoix(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "choix", nil)
}

func ForumPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "form", nil)
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func ListPage(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("../perso.json")
	if err != nil {
		fmt.Println("Erreur dans la lecture du json : ", err)
		http.Error(w, "Erreur dans la lecture du JSON", http.StatusInternalServerError)
		return
	}

	var allArticles []backend.PersoData

	err = json.Unmarshal(content, &allArticles)
	if err != nil {
		fmt.Println("Erreur > ", err.Error())
		http.Error(w, "Erreur lors de la désérialisation du JSON", http.StatusInternalServerError)
		return
	}

	templates.Temp.ExecuteTemplate(w, "list", allArticles)
}

func InitSexeHomme(w http.ResponseWriter, r *http.Request) {
	sexe = "Homme"
	img = "/static/img/homme.png"
	http.Redirect(w, r, "/form", http.StatusMovedPermanently)
}

func InitSexeFemme(w http.ResponseWriter, r *http.Request) {
	sexe = "Femme"
	img = "/static/img/femme.png"
	http.Redirect(w, r, "/form", http.StatusMovedPermanently)
}
