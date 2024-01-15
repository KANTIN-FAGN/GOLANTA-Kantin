package controller

import (
	"encoding/json"
	"fmt"
	"golantah/backend"
	"golantah/temp"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// ArticlePage est la fonction handler de la page d'article qui permet d'afficher un article
func ArticlePage(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	articleID, err := strconv.Atoi(queryID)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	var jsonData backend.JSONData

	jsonDataFile, err := ioutil.ReadFile("perso.json")
	err = json.Unmarshal(jsonDataFile, &jsonData)

	var foundArticle *backend.Article
	for _, category := range jsonData.Categories {
		for _, article := range category.Articles {
			if article.Id == articleID {
				foundArticle = &article
				break
			}
		}
		if foundArticle != nil {
			break
		}
	}

	if foundArticle == nil {
		templates.Temp.ExecuteTemplate(w, "erreur", nil)
	}


	templates.Temp.ExecuteTemplate(w, "article", r)
}

// IndexPage est la fonction handel de la page d'accueil
func IndexPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "index", r)
}


// RecuDatas permet de récupérer les données de l'ajout d'un perso et de les traiter
func RecuDatas(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	nom := r.FormValue("nom")
	prenom := r.FormValue("prenom")
	fort := r.FormValue("point-fort")
	faible := r.FormValue("point-faible")
	fmt.Println(nom)
	fmt.Println(prenom)

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filepath := "./assets/img/" + handler.Filename
	f, _ := os.Create(filepath)
	defer f.Close()
	io.Copy(f, file)

	jsonFile, _ := ioutil.ReadFile("perso.json")
	var jsonData backend.JSONData
	json.Unmarshal(jsonFile, &jsonData)

	ids, _ := backend.GetArticleIDs("perso.json")

	rand.Seed(time.Now().UnixNano())
	var newID int

	for {
		newID = rand.Intn(8999) + 1000
		if !backend.IsIDPresent(newID, ids) {
			break
		}
	}

	nouveauPerso := backend.Perso{
		Id:     newID,
		Nom:  	nom,
		Prenom: prenom,
		Fort: 	fort,
		Faible: faible,
		Image:  handler.Filename,
	}

	backend.AddPerso(&jsonData, nouveauPerso)

	updatedData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Println("Erreur en convertissant les données en JSON :", err)
		return
	}

	ioutil.WriteFile("perso.json", updatedData, 0644)

	http.Redirect(w, r, "/new_perso", http.StatusSeeOther)
}
