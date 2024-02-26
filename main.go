package main

import (
	"bekabar_chat/config"
	"bekabar_chat/router"
	"bekabar_chat/structs"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gorm.io/gorm"
)

func MyNewService(
	DBConnection *gorm.DB,
) *structs.MyConnectionService {
	return &structs.MyConnectionService{DB: DBConnection}
}

func main() {
	config.LoadEnvVariable()

	port := os.Getenv("PORT")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// set up DB connection here
	DBConnection := config.InitDatabase()

	services := MyNewService(DBConnection)

	// all app routes here
	appRoute := router.Dispatch(services)

	srv := &http.Server{
		Addr: "localhost:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      appRoute, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		fmt.Println("⚡️ [" + os.Getenv("APPLICATION_ENV") + "] - " + os.Getenv("APP_NAME") + " IS RUNNING ON PORT - " + "http://localhost:" + port)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	fmt.Println("shutting down")
	os.Exit(0)

}
