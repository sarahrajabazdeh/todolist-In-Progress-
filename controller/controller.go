package controller

import (
	"github.com/sarahrajabazdeh/todolist/dataservice"
)

type ControllerInterface interface {
}
type HttpController struct {
	DS dataservice.DataserviceInterface
}
