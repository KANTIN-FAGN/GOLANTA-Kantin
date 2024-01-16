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
	http.HandleFunc("/choix", controller.DisplayChoix)
	http.HandleFunc("/treatment/homme", controller.InitSexeHomme)
	http.HandleFunc("/treatment/femme", controller.InitSexeFemme)
	http.HandleFunc("/form", controller.ForumPage)
	http.HandleFunc("/list", controller.ListPage)

	http.HandleFunc("/", controller.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
