package routeur

import (
	"fmt"
	"golantah/backend"
	controller "golantah/controller"
	"log"
	"net/http"
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
	http.HandleFunc("/delete", backend.DeletePage)
	http.HandleFunc("/submit", controller.RecuDatas)
	http.HandleFunc("/modif", backend.DisplayModif)
	http.HandleFunc("/modif/treatment", backend.TreatmentModif)
	http.HandleFunc("/treatment/equipe/verte", controller.InitEquipeVerte)
	http.HandleFunc("/treatment/equipe/jaune", controller.InitEquipeJaune)
	http.HandleFunc("/treatment/equipe/rouge", controller.InitEquipeRouge)
	http.HandleFunc("/treatment/equipe/bleu", controller.InitEquipeBleu)
	http.HandleFunc("/choix/equipe", controller.DisplayChoixEquipe)

	http.HandleFunc("/", controller.DefaultHandler)

	// Démarrage du serveur
	log.Println("[✅] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8080/accueil")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
