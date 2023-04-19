package handler

import (
	"openaigo/config"
	"openaigo/src/handler/seedhandler"
	"openaigo/src/lib/logger"
	"openaigo/src/service"
	"openaigo/src/tools/validator"
)

type ICore interface {
	Seed() seedhandler.IHandler
}

type handler struct {
	seed seedhandler.IHandler
}

func New(
	cfg config.IConfig,
	serviceCore service.ICore,
	v validator.IValidate,
	l logger.ILogger,
) ICore {
	return &handler{
		seed: seedhandler.New(cfg, serviceCore, v, l),
	}
}

func (h *handler) Seed() seedhandler.IHandler {
	return h.seed
}
