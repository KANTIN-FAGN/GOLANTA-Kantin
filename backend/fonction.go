package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
// GetArticleIDs récupère tous les id des perso du fichier perso.json
func GetArticleIDs(data []PersoData) ([]int, error) {
    var ids []int
    for _, entry := range data {
        id, err := strconv.Atoi(entry.ID)
        if err != nil {
            return nil, err
        }
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
		file, err := ioutil.ReadFile("../perso.json")
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
		err = ioutil.WriteFile("../perso.json", newData, 0644)
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
	jsonFile, err := ioutil.ReadFile("../perso.json")
	if err != nil {
		fmt.Println("Erreur en lisant le fichier JSON :", err)
		http.Error(w, "Erreur en lisant le fichier JSON", http.StatusInternalServerError)
		return
	}

	// Désérialiser le JSON dans une slice de PersoData
	var jsonData []PersoData
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Println("Erreur en désérialisant les données JSON :", err)
		http.Error(w, "Erreur en désérialisant les données JSON", http.StatusInternalServerError)
		return
	}

	// Générer un nouvel ID unique
	rand.Seed(time.Now().UnixNano())
	var newID int

	ids, err := GetArticleIDs(jsonData)
	if err != nil {
    	fmt.Println("Erreur en obtenant les IDs :", err)
    	http.Error(w, "Erreur en obtenant les IDs", http.StatusInternalServerError)
    	return
	}

	maxAttempts := 100
	for attempt := 0; attempt < maxAttempts; attempt++ {
    	newID = rand.Intn(8999) + 1000
    	if !IsIDPresent(newID, ids) {
       	 break
    	}
	}

	// Créer un nouvel article avec le nouvel ID
	nouvelArticle := PersoData{
		ID:     strconv.Itoa(newID),
		Nom:    nom,
		Prenom: prenom,
		Age:    age,
		Taille: taille,
		Poids:  poids,
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
	err = ioutil.WriteFile("../perso.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Erreur en écrivant dans le fichier JSON :", err)
		http.Error(w, "Erreur en écrivant dans le fichier JSON", http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page de liste après l'ajout
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}


func AddPerso(jsonData *[]PersoData, perso PersoData) error {
	*jsonData = append(*jsonData, perso)
	return nil
}

func GetAllPerso() []PersoData {
	var allArticles []PersoData
	return allArticles
}