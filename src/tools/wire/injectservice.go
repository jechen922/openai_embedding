package wire

import (
	"encoding/json"
	"openaigo/config"
	"openaigo/src/database"
	"openaigo/src/database/postgres"
	"openaigo/src/handler"
	"openaigo/src/lib/logger"
	"openaigo/src/repository"
	"openaigo/src/router"
	"openaigo/src/service"
	"openaigo/src/tools/validator"
	"openaigo/src/utils/openai"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet( //nolint:deadcode,unused,varcheck
	provideRepository,
	provideService,
	provideHandler,
	provideRouter,
	provideFiber,
	provideDatabase,
	provideOpenAI,
	provideValidator,
	provideLogger,
)

func provideRepository() repository.ICore {
	return repository.New()
}

func provideService(db database.IDatabase, repo repository.ICore, ai openai.IAI) service.ICore {
	return service.New(db, repo, ai)
}

func provideHandler(
	cfg config.IConfig,
	s service.ICore,
	v validator.IValidate,
	l logger.ILogger,
) handler.ICore {
	return handler.New(cfg, s, v, l)
}

func provideRouter(h handler.ICore) router.IRouter {
	return router.New(h)
}

func provideFiber() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
	})
}

func provideDatabase() (database.IDatabase, error) {
	postgresDB, err := postgres.New()
	return database.New(postgresDB), err
}

func provideOpenAI(cfg config.IConfig) openai.IAI {
	return openai.New(cfg)
}

func provideValidator(l logger.ILogger) validator.IValidate {
	return validator.New(l)
}

func provideLogger(cfg config.IConfig) logger.ILogger {
	return logger.New(cfg.GetLoggerENV())
}
