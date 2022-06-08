package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sarahrajabazdeh/todolist/config"
	"github.com/sarahrajabazdeh/todolist/controller"
	"github.com/sarahrajabazdeh/todolist/database"
	"github.com/sarahrajabazdeh/todolist/dataservice"
	"github.com/sarahrajabazdeh/todolist/router"
)

func main() {
	log.Println("Stating Go server")

	// read the configuration
	config.ReadConfig()

	// create the db connection
	db := database.CreateDatabase()

	ds := dataservice.Dataservice{
		DB: &db,
	}

	ctrl := controller.HttpController{
		DS: &ds,
	}

	r := mux.NewRouter()
	router.SetupRoutes(r, &ctrl)

	err := http.ListenAndServe(":"+config.Config.Server.Port, cors.AllowAll().Handler(r))
	if err != nil {
		log.Println(err)
	}

}
