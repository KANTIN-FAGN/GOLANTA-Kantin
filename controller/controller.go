package controller

import (
	templates "golantah/temp"
	"net/http"
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

