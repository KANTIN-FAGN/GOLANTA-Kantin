package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"golantah/backend"
	"golantah/temp"
)

// AddArticlePage est la fonction handler de la page d'ajout d'articles
func AddArticlePage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "newarticle", nil)
}

// DeletePage est la fonction qui permet de supprimer un article de la base de donnée
func DeletePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		articleID := r.FormValue("perso_id")

		id, err := strconv.Atoi(articleID)
		if err != nil {
			http.Error(w, "Invalid article ID", http.StatusBadRequest)
			return
		}

		file, _ := ioutil.ReadFile("perso.json")

		var jsonData backend.JSONData
		json.Unmarshal(file, &jsonData)

		found := false
		for i := range jsonData.Categories {
			for j, article := range jsonData.Categories[i].Articles {
				if article.Id == id {
					jsonData.Categories[i].Articles = append(jsonData.Categories[i].Articles[:j], jsonData.Categories[i].Articles[j+1:]...)
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		newData, _ := json.MarshalIndent(jsonData, "", "  ")
		ioutil.WriteFile("perso.json", newData, 0644)

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

// DefaultHandler est la fonction qui redirige vers la page 404 en cas de route inconnue
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	err := templates.Temp.ExecuteTemplate(w, "erreur", nil)
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}
