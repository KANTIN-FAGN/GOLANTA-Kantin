package main

import (
	"fmt"
	"golantah/backend"
	"golantah/routeur"
	"golantah/temp"
)

func main() {
	active, user := backend.CheckRememberStatus("remember.json")
	if active {
		fmt.Println("Une session a été sauvegardée")
		backend.GlobalSession = backend.Session{Username: user, State: backend.GetAccountState(user), Mail: backend.GetAccountMail(user)}
		fmt.Println("Session initialisée")
	}
	templates.InitTemplate()
	routeur.Initserv()
}
