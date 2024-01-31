package router

import (
	"bekabar_chat/apps/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"bekabar_chat/apps/user"
)

func Dispatch(DBConnection *gorm.DB) *mux.Router {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.AppResponse(w, http.StatusOK, "Hello World")
	}).Methods("GET")

	myRouter.Handle("/register", user.RegisterUser(DBConnection)).Methods("POST")

	return myRouter
}
