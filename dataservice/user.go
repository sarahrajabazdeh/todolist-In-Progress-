package dataservice

import "github.com/sarahrajabazdeh/todolist/model"

func (ds *Dataservice) GetUsers() []model.User {

	return ds.DB.GetUsers()
}
