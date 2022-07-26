package dataservice

import (
	"github.com/sarahrajabazdeh/todolist/database"
	"github.com/sarahrajabazdeh/todolist/model"
)

// package that contains all the business logic

type DataserviceInterface interface {
	GetUsers() []model.User
}

type Dataservice struct {
	DB database.DatabaseInterface
}
