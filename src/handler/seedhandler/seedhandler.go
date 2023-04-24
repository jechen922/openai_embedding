package seedhandler

import (
	"github.com/gofiber/fiber/v2"
	"openaigo/config"
	"openaigo/openai/embedding"
	"openaigo/src/lib/file"
	"openaigo/src/lib/logger"
	"openaigo/src/service"
	"openaigo/src/tools/validator"
)

type IHandler interface {
	Train(ctx *fiber.Ctx) error
}

type seedHandler struct {
	config    config.IConfig
	service   service.ICore
	validator validator.IValidate
	logger    logger.ILogger
}

func New(cfg config.IConfig, serviceCore service.ICore, v validator.IValidate, iLog logger.ILogger) IHandler {
	return &seedHandler{config: cfg, service: serviceCore, validator: v, logger: iLog}
}

func (h *seedHandler) Train(ctx *fiber.Ctx) error {
	csvRecords := file.ReadCSVByFields("./resources/yile/yile.csv", "title", "heading", "content")
	sections := make([]embedding.Section, 0, len(csvRecords))
	for _, record := range csvRecords {
		if record["title"] != "員工福利規範" {
			continue
		}
		sections = append(sections, embedding.Section{
			Title:   record["title"],
			Heading: record["heading"],
			Content: record["content"],
		})
	}

	if err := h.service.Seed().SaveSections(sections); err != nil {
		return err
	}
	return ctx.Send([]byte("OK!"))
}
