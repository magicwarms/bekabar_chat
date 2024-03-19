package main

import (
	"bekabar_chat/config"
	"bekabar_chat/di"
	"bekabar_chat/server"
	"log"
	"os"
)

func main() {
	config.LoadEnvVariable()

	if err := RunApplication(os.Getenv("APPLICATION_ENV")); err != nil {
		log.Fatal(err)
	}

}

func RunApplication(env string) error {
	buildContainers := di.BuildContainer(env)

	return buildContainers.Invoke(func(apiServer *server.APIServer) error {
		return apiServer.Start()
	})
}
