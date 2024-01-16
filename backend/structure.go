package backend

type PersoData struct {
	ID      string `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Age     string `json:"age"`
	Taille  int    `json:"taille"`
	Poids   int    `json:"poids"`
	Sexe    string `json:"sexe"`
	Image   string `json:"image"`
}
