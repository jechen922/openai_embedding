package router

import (
	"openaigo/config"
	"openaigo/src/handler"
	
	"github.com/gofiber/fiber/v2"
)

type (
	IRouter interface {
		Set(r fiber.Router)
	}

	Router struct {
		config  config.Config
		handler handler.ICore
	}
)

func New(h handler.ICore) IRouter {
	return Router{handler: h}
}

func (r Router) Set(fRoute fiber.Router) {
	fRoute.Get("/train", r.handler.Seed().Train)

}
