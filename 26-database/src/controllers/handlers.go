package controllers

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

	"database-golang/src/services"
	"database-golang/src/types"

	"github.com/gorilla/mux"
)

func usuariosPOST(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("failed on reading request body"))
		return
	}

	var usuario types.User
	if err = json.Unmarshal(payload, &usuario); err != nil {
		w.Write([]byte("error parsing json into struct"))
		return
	}

	log.Println(usuario)

	createdUser, err := services.CreateUser(&usuario)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user successfully inserted, id: %d", createdUser.ID)))
}

func usuariosGET(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsuarios()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("error on converting users to json"))
		return
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func CreateHandlers(router *mux.Router) *mux.Router {
	router.Use(commonMiddleware)
	router.HandleFunc("/usuarios", usuariosPOST).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", usuariosGET).Methods(http.MethodGet)
	return router
}
