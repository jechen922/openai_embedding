package router

import (
	"github.com/gofiber/fiber/v2"
	"openaigo/src/handler/seedhandler"
)

type (
	IRouter interface {
		Set(r fiber.Router)
	}

	//Router struct {
	//	config  config.Config
	//	handler handler.ICore
	//}
)

//func New(h handler.ICore) IRouter {
//	return Router{handler: h}
//}

//func (r Router) Set(fRoute fiber.Router) {

func Set(fRoute fiber.Router) {
	fRoute.Get("/train", seedhandler.Seed)

}
