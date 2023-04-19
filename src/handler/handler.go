package handler

import (
	"openaigo/config"
	"openaigo/src/handler/customerservicehandler"
	"openaigo/src/handler/seedhandler"
	"openaigo/src/lib/logger"
	"openaigo/src/service"
	"openaigo/src/tools/validator"
)

type ICore interface {
	Seed() seedhandler.IHandler
	CustomerService() customerservicehandler.IHandler
}

type handler struct {
	seed            seedhandler.IHandler
	customerService customerservicehandler.IHandler
}

func New(
	cfg config.IConfig,
	serviceCore service.ICore,
	v validator.IValidate,
	l logger.ILogger,
) ICore {
	return &handler{
		seed:            seedhandler.New(cfg, serviceCore, v, l),
		customerService: customerservicehandler.New(cfg, serviceCore, v, l),
	}
}

func (h *handler) Seed() seedhandler.IHandler {
	return h.seed
}

func (h *handler) CustomerService() customerservicehandler.IHandler {
	return h.customerService
}
