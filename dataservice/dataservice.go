package dataservice

import (
	"github.com/sarahrajabazdeh/todolist/database"
)

// package that contains all the business logic

type DataserviceInterface interface {
}

type Dataservice struct {
	DB database.DatabaseInterface
}
