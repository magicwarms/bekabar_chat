package di

import (
	"bekabar_chat/apps/user"
	"bekabar_chat/config"
	"bekabar_chat/server"

	"go.uber.org/dig"
)

func BuildContainer(env string) *dig.Container {
	dryRun := false
	if env == "testing" {
		dryRun = true
	}
	container := dig.New(dig.DryRun(dryRun))

	container.Provide(config.InitDatabase)
	container.Provide(user.NewUserService)
	container.Provide(user.NewUserHandler)

	container.Provide(server.NewAPIServer)

	return container
}
