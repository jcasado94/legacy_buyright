package server

import (
	"encoding/json"
	"errors"
	"gobuyright/pkg/entity"
	"net/http"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService entity.UserService
}

func NewUserRouter(u entity.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}

	router.HandleFunc("/", userRouter.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")

	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJson(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ur.userService.GetByUsername(username)
	if err != nil {
		WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	WriteJson(w, http.StatusOK, user)
}

func decodeUser(r *http.Request) (entity.IUser, error) {
	var u entity.IUser
	if r.Body == nil {
		return u, errors.New("No request body")
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)

	return u, err
}
