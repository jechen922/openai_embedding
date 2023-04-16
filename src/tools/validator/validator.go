package validator

import (
	"openaigo/src/lib/logger"
	
	"github.com/gofiber/fiber/v2"
)

type IValidate interface {
	Parse(ctx *fiber.Ctx, st interface{}) error
}

func New(l logger.ILogger) IValidate {
	return newCustomValidate(l)
}
