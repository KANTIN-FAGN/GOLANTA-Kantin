package backend

type PersoData struct {
	ID     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Age    string `json:"age"`
	Taille string `json:"taille"`
	Poids  string `json:"poids"`
	Sexe   string `json:"sexe"`
	Image  string `json:"image"`
}
