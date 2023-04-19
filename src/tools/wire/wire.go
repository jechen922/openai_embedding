//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gofiber/fiber/v2"
	"openaigo/config"
	"openaigo/src/router"

	"github.com/google/wire"
)

type Application struct {
	Fiber  *fiber.App
	Router router.IRouter
}

func newApplication(
	fiberApp *fiber.App,
	r router.IRouter,
) *Application {
	return &Application{
		Fiber:  fiberApp,
		Router: r,
	}
}

func InitializeApplication(cfg config.IConfig) (*Application, error) {
	wire.Build(
		serviceSet,
		newApplication,
	)
	return &Application{}, nil
}
