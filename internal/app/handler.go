package app

import (
	"github.com/dshepett/task-manager/pkg/responder"
	"net/http"
)

func (a *App) addRoutes() {
	a.router.HandleFunc("/", index)
	a.router.HandleFunc("/users", a.showUsersHandler).Methods("GET")
	a.router.HandleFunc("/users/add", a.addUserHandler).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.showUserHandler).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/edit", a.editUserHandler).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}/delete", a.deleteUserHandler).Methods("DELETE")
}

func index(w http.ResponseWriter, r *http.Request) {
	responder.RespondWithJson(w, 200, map[string]bool{"alive": true})
}
