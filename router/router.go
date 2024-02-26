package router

import (
	"bekabar_chat/apps/user"
	"bekabar_chat/apps/utils"
	"bekabar_chat/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func Dispatch(services *structs.MyConnectionService) *mux.Router {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.AppResponse(w, http.StatusOK, "Hello World")
	}).Methods("GET")

	myRouter.Handle("/register", user.RegisterUser(services)).Methods(http.MethodPost, http.MethodOptions)
	myRouter.Use(mux.CORSMethodMiddleware(myRouter))

	return myRouter
}
