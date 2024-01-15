package routeur

import (
	"fmt"
	"log"
	"net/http"
	"golantah/controller"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.IndexPage)
	

	http.HandleFunc("/article", controller.ArticlePage)
	http.HandleFunc("/new_article", controller.AddArticlePage)
	http.HandleFunc("/submit", controller.RecuDatas)
	http.HandleFunc("/delete", controller.DeletePage)



	http.HandleFunc("/", controller.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
