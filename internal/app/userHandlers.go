package app

import (
	"github.com/dshepett/task-manager/internal/models"
	"github.com/dshepett/task-manager/pkg/responder"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (a *App) addUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "error during reading form's data")
		return
	}
	name := r.Form.Get("name")
	user := &models.User{Name: name}
	if err := a.storage.UserRepository.Create(user); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "could not add user")
		return
	}
	responder.RespondWithJson(w, http.StatusOK, user)
}

func (a *App) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := &models.User{Id: id}
	if err := a.storage.UserRepository.GetById(user); err != nil {
		responder.RespondWithError(w, http.StatusNotFound, "could not find user with such id")
		return
	}
	responder.RespondWithJson(w, http.StatusOK, user)
}

func (a *App) showUsersHandler(w http.ResponseWriter, r *http.Request) {
	if users, err := a.storage.UserRepository.GetAll(); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "could not get users")
	} else {
		responder.RespondWithJson(w, http.StatusOK, users)
	}
}

func (a *App) editUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := r.ParseForm(); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "error during reading form's data")
		return
	}
	name := r.Form.Get("name")
	user := &models.User{Id: id, Name: name}
	if err := a.storage.UserRepository.Edit(user); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "could not edit user")
	} else {
		responder.RespondWithJson(w, http.StatusOK, user)
	}

}

func (a *App) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := a.storage.UserRepository.Delete(id); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "could not delete user")
	} else {
		responder.RespondWithJson(w, http.StatusOK, map[string]bool{"success": true})
	}
}
