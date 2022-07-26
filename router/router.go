package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahrajabazdeh/todolist/controller"
)

var ctrl controller.ControllerInterface

func SetupRoutes(router *mux.Router, c controller.ControllerInterface) {
	ctrl = c

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	// GetUsers return all user in database path "localhost:9090/users" GET
	router.HandleFunc("/users", ctrl.GetUsers).Methods("GET")

}
