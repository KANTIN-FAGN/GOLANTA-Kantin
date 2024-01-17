package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// IsIDPresent vérifie si un id est présent dans une liste d'id
func IsIDPresent(id int, ids []int) bool {
	for _, existingID := range ids {
		if existingID == id {
			return true
		}
	}
	return false
}

// GetArticleIDs récupère tous les id des perso du fichier perso.json
func GetArticleIDs(filename string) ([]int, error) {
	var jsonData []PersoData


	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Erreur de lecture du fichier : %v", err)
	}

	if err := json.Unmarshal(raw, &jsonData); err != nil {
		return nil, fmt.Errorf("Erreur de désérialisation JSON : %v", err)
	}

	var articleIDs []int
	for _, personne := range jsonData {
		id, err := strconv.Atoi(personne.ID)
		if err != nil {
			return nil, fmt.Errorf("Erreur de conversion de l'ID en entier : %v", err)
		}

		articleIDs = append(articleIDs, id)
	}

	return articleIDs, nil
}

// DeletePage est la fonction qui permet de supprimer un article de la base de donnée
func DeletePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Récupérer l'ID de la personne à supprimer depuis le formulaire
		idStr := r.FormValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Lire le fichier JSON existant
		file, err := ioutil.ReadFile("perso.json")
		if err != nil {
			http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
			return
		}

		// Désérialiser le JSON dans la slice de PersoData
		var jsonData []PersoData
		err = json.Unmarshal(file, &jsonData)
		if err != nil {
			http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
			return
		}

		// Rechercher la personne dans la slice de données et la supprimer si elle est trouvée
		found := false
		for j, personne := range jsonData {
			personneID, err := strconv.Atoi(personne.ID)
			if err != nil {
				http.Error(w, "Error converting person ID to integer", http.StatusInternalServerError)
				return
			}

			if personneID == id {
				// Supprimer la personne de la slice
				jsonData = append(jsonData[:j], jsonData[j+1:]...)
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Person not found", http.StatusNotFound)
			return
		}

		// Sérialiser la slice de données mise à jour en JSON avec indentation
		newData, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			http.Error(w, "Error encoding JSON data", http.StatusInternalServerError)
			return
		}

		// Écrire les données mises à jour dans le fichier JSON
		err = ioutil.WriteFile("perso.json", newData, 0644)
		if err != nil {
			http.Error(w, "Error writing to JSON file", http.StatusInternalServerError)
			return
		}

		// Rediriger vers la page de liste après la suppression
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	}
}
