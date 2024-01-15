package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
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

// GetArticleIDs récupère tous les id des articles du fichier blog.json
func GetArticleIDs(filename string) ([]int, error) {
	var data map[string]interface{}

	raw, _ := ioutil.ReadFile(filename)

	json.Unmarshal(raw, &data)

	var articleIDs []int
	categories, ok := data["categories"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("Champ 'categories' non trouvé ou incorrect")
	}

	for _, category := range categories {
		cat, ok := category.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Structure de catégorie incorrecte")
		}

		articles, ok := cat["articles"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("Champ 'articles' non trouvé ou incorrect")
		}

		for _, article := range articles {
			art, ok := article.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("Structure d'article incorrecte")
			}

			id, ok := art["id"].(float64)
			if !ok {
				return nil, fmt.Errorf("Champ 'id' non trouvé ou incorrect")
			}

			articleIDs = append(articleIDs, int(id))
		}
	}

	return articleIDs, nil
}

// TitleContains vérifie si une chaine de caractère substr est contenue dans une chaine s
func TitleContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// AddArticle prend en argument une structure json, un nom de catégorie et un article puis ajoute l'article dans la bonne catégorie
func AddPerso(jsonData *JSONData, categoryName string, article Article) error {
	for i := range jsonData.Categories {
		if jsonData.Categories[i].Name == categoryName {
			jsonData.Categories[i].Articles = append(jsonData.Categories[i].Articles, article)
			return nil
		}
	}
	return fmt.Errorf("category '%s' not found", categoryName)
}

// GetAllArticles récupère tous les articles contenus dans une structure JSONData
func GetAllArticles(jsonData JSONData) []Article {
	var allArticles []Article
	for _, categorie := range jsonData.Categories {
		allArticles = append(allArticles, categorie.Articles...)
	}

	return allArticles
}
