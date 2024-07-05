// go:build wireinject
//+build wireinject

package main

import (
	"gallery_go/app"
	"gallery_go/controller"
	"gallery_go/repository"
	"gallery_go/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

func initializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		userSet,
		app.NewRouter,
		NewServer,
	)

	return nil
}

