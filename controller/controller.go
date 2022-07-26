package controller

import (
	"net/http"

	"github.com/sarahrajabazdeh/todolist/dataservice"
)

type ControllerInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type HttpController struct {
	DS dataservice.DataserviceInterface
}
