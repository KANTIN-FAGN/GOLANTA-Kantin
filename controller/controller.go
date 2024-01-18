package controller

import (
	"encoding/json"
	"fmt"
	"golantah/backend"
	templates "golantah/temp"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var Sexe string
var Img string

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

func ListPage(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("perso.json")
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
	Sexe = "Homme"
	Img = "/static/img/site/homme.png"
	http.Redirect(w, r, "/form", http.StatusMovedPermanently)
}

func InitSexeFemme(w http.ResponseWriter, r *http.Request) {
	Sexe = "Femme"
	Img = "/static/img/site/femme.png"
	http.Redirect(w, r, "/form", http.StatusMovedPermanently)
}

func RecuDatas(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	fmt.Println("here")

	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	age := r.FormValue("age")
	taille := r.FormValue("taille")
	poids := r.FormValue("poids") // Correction du nom du champ

	fmt.Println(nom)
	fmt.Println(prenom)

	// Lire le fichier JSON existant
	jsonFile, err := os.ReadFile("perso.json")
	if err != nil {
		fmt.Println("Erreur en lisant le fichier JSON :", err)
		http.Error(w, "Erreur en lisant le fichier JSON", http.StatusInternalServerError)
		return
	}

	// Désérialiser le JSON dans une slice de PersoData
	var jsonData []backend.PersoData
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Println("Erreur en désérialisant les données JSON :", err)
		http.Error(w, "Erreur en désérialisant les données JSON", http.StatusInternalServerError)
		return
	}

	// Générer un nouvel ID unique
	rand.Seed(time.Now().UnixNano())
	var newID int

	ids, err := backend.GetArticleIDs(jsonData)
	if err != nil {
		fmt.Println("Erreur en obtenant les IDs :", err)
		http.Error(w, "Erreur en obtenant les IDs", http.StatusInternalServerError)
		return
	}

	maxAttempts := 100
	for attempt := 0; attempt < maxAttempts; attempt++ {
		newID = rand.Intn(8999) + 1000
		if !backend.IsIDPresent(newID, ids) {
			break
		}
	}

	// Créer un nouvel article avec le nouvel ID
	nouvelArticle := backend.PersoData{
		ID:     newID,
		Nom:    nom,
		Prenom: prenom,
		Age:    age,
		Taille: taille,
		Poids:  poids,
		Sexe: Sexe,
		Image: Img,

	}

	// Ajouter le nouvel article à la slice de données
	jsonData = append(jsonData, nouvelArticle)

	// Sérialiser la slice de données mise à jour en JSON avec indentation
	updatedData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Erreur en convertissant les données en JSON :", err)
		http.Error(w, "Erreur en convertissant les données en JSON", http.StatusInternalServerError)
		return
	}

	// Écrire les données mises à jour dans le fichier JSON
	err = os.WriteFile("perso.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur en écrivant dans le fichier JSON :", err)
		http.Error(w, "Erreur en écrivant dans le fichier JSON", http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page de liste après l'ajout
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}