package controller

import (
	"encoding/json"
	"net/http"
)

func (ctrl *HttpController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users := ctrl.DS.GetUsers()

	// respond to FE
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
