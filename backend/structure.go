package backend

// Article est une structure qui stock toutes les données d'un article
type Perso struct {
	Id     int    `json:"id"`
	Nom  string `json:"nom"`
	Prenom  string `json:"prenom"`
	Image  string `json:"image"`
	Fort string `json:"fort"`
	Faible   string `json:"faible"`
}

// IndexData est une structure qui gère les données envoyées à la page index
type IndexData struct {
	Articles   []Perso
}

// ArticleData est une structure qui gère les données envoyées à la page index
type ArticleData struct {
	Data       map[string]interface{}
}
