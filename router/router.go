package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarahrajabazdeh/todolist/controller"
)

func SetupRoutes(router *mux.Router, ctrl controller.ControllerInterface) {

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

}
