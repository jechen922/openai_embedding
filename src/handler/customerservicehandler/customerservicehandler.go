package customerservicehandler

import (
	"github.com/gofiber/fiber/v2"
	"openaigo/config"
	"openaigo/src/lib/logger"
	"openaigo/src/service"
	"openaigo/src/tools/validator"
)

type IHandler interface {
	Question(ctx *fiber.Ctx) error
}

type customerServiceHandler struct {
	config    config.IConfig
	service   service.ICore
	validator validator.IValidate
	logger    logger.ILogger
}

func New(cfg config.IConfig, serviceCore service.ICore, v validator.IValidate, iLog logger.ILogger) IHandler {
	return &customerServiceHandler{
		config:    cfg,
		service:   serviceCore,
		validator: v,
		logger:    iLog,
	}
}

func (h *customerServiceHandler) Question(ctx *fiber.Ctx) error {
	type question struct {
		Question string `json:"Question"`
	}
	args := question{}
	if err := h.validator.Parse(ctx, &args); err != nil {
		return err
	}

	ans, err := h.service.CustomerService().Ask(args.Question)
	if err != nil {
		return err
	}
	return ctx.JSON(map[string]string{
		"Answer": ans,
	})
}
