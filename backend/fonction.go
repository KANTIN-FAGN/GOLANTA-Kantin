package backend

import (
	"encoding/json"
	"fmt"
	templates "golantah/temp"
	"net/http"
	"os"
	"strconv"
)

// GetArticleIDs récupère tous les id des perso du fichier perso.json
func GetArticleIDs(data []PersoData) ([]int, error) {
	var ids []int
	for _, entry := range data {
		id := entry.ID
		ids = append(ids, id)
	}
	return ids, nil
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
		file, err := os.ReadFile("perso.json")
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
			if personne.ID == id {
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
		err = os.WriteFile("perso.json", newData, 0644)
		if err != nil {
			http.Error(w, "Error writing to JSON file", http.StatusInternalServerError)
			return
		}

		// Rediriger vers la page de liste après la suppression
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	}
}

func IsIDPresent(id int, ids []int) bool {
	for _, existingID := range ids {
		if existingID == id {
			return true
		}
	}
	return false
}


func AddPerso(jsonData *[]PersoData, perso PersoData) error {
	*jsonData = append(*jsonData, perso)
	return nil
}

func GetAllPerso() []PersoData {
	var allArticles []PersoData
	return allArticles
}

func DisplayModif(w http.ResponseWriter, r *http.Request) {
	var tosend PersoData

	id, err := strconv.Atoi(r.URL.Query().Get("char"))
	if err != nil {
		return
	}

	data, err := os.ReadFile("perso.json")
	if err != nil {
		return
	}

	var persodata []PersoData
	err = json.Unmarshal(data, &persodata)
	if err != nil {
		return
	}

	for _, i := range persodata {
		if i.ID == id {
			tosend = i
		}
	}
	templates.Temp.ExecuteTemplate(w, "modif", tosend)
}

func TreatmentModif(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Println("Erreur pour ID")

	}

	data, err := os.ReadFile("perso.json")
	if err != nil {
		fmt.Println("Erreur dans la lecture du JSON")
	}

	var persodata []PersoData
	err = json.Unmarshal(data, &persodata)
	if err != nil {
		fmt.Println("Erreur pour recupération des données du JSON")
	}
	var indexPerso int
	for i, perso := range persodata {
		if perso.ID == id {
			indexPerso = i
			break
		}
	}

	persodata[indexPerso].Nom = r.FormValue("nom")
	persodata[indexPerso].Prenom = r.FormValue("prenom")
	persodata[indexPerso].Age = r.FormValue("age")
	persodata[indexPerso].Taille = r.FormValue("taille")
	persodata[indexPerso].Poids = r.FormValue("poids")

	
	updatedData, err := json.Marshal(persodata)
	if err != nil {
		fmt.Println("Erreur en convertissant les données en JSON :", err)
		http.Error(w, "Erreur en convertissant les données en JSON", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("perso.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur en écrivant dans le fichier JSON :", err)
		http.Error(w, "Erreur en écrivant dans le fichier JSON", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusMovedPermanently)
}
