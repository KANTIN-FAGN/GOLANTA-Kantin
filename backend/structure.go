package backend

type PersoData struct {
	ID      string `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Age     string `json:"age"`
	Force   int    `json:"force"`
	Agility int    `json:"agility"`
	Taille  int    `json:"taille"`
	Poids   int    `json:"poids"`
	Sexe    string `json:"sexe"`
	Image   string `json:"image"`
}
