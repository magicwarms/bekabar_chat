package server

import (
	"bekabar_chat/apps/user"
	"bekabar_chat/apps/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(
	userHandler *user.UserHandler,
) *APIServer {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.AppResponse(w, http.StatusOK, "Hello World")
	}).Methods("GET")

	myRouter.Handle("/register", userHandler.RegisterUser()).Methods(http.MethodPost, http.MethodOptions)
	myRouter.Handle("/all", userHandler.GetAllUser()).Methods(http.MethodGet, http.MethodOptions)

	myRouter.Use(mux.CORSMethodMiddleware(myRouter))

	return &APIServer{
		httpServer: &http.Server{
			Addr: "localhost:" + os.Getenv("PORT"),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      myRouter, // Pass our instance of gorilla/mux in.
		},
	}
}

func (s *APIServer) Start() error {

	fmt.Println("⚡️ [" + os.Getenv("APPLICATION_ENV") + "] - " + os.Getenv("APP_NAME") + " IS RUNNING ON PORT - " + "http://localhost:" + os.Getenv("PORT"))

	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
