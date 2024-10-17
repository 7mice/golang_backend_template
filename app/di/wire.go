// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"ginTemplate/app/connections"
	"ginTemplate/internal/controller"
	"ginTemplate/internal/repository"
	"ginTemplate/internal/service"
	"github.com/google/wire"
)

var connectionsSet = wire.NewSet(
	connections.ConnectToDB,
	connections.ConnectToNatsBroker,
	connections.ConnectToRedis,
)

var testSet = wire.NewSet(
	repository.TestRepositoryInit,
	wire.Bind(new(repository.TestRepository), new(*repository.TestRepositoryImpl)),
	service.TestServiceInit,
	wire.Bind(new(service.TestService), new(*service.TestServiceImpl)),
	controller.TestControllerInit,
	wire.Bind(new(controller.TestController), new(*controller.TestControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization,
		connectionsSet,
		testSet)
	return nil
}
