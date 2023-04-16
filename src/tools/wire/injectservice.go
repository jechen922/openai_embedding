package wire

import (
	"database/sql"
	"encoding/json"
	"openaigo/config"
	"openaigo/src/database"
	"openaigo/src/handler"
	"openaigo/src/lib/logger"
	"openaigo/src/repository"
	"openaigo/src/router"
	"openaigo/src/service"
	"openaigo/src/tools/validator"

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
	provideValidator,
	provideLogger,
)

func provideRepository() repository.ICore {
	return repository.New()
}

func provideService(db database.IDatabase) service.ICore {
	return service.New(db)
}

func provideHandler(
	cfg config.Config,
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

func provideDatabase(postgresDB *sql.DB) database.IDatabase {
	return database.New(postgresDB)
}

func provideValidator(l logger.ILogger) validator.IValidate {
	return validator.New(l)
}

func provideLogger(cfg config.Config) logger.ILogger {
	return logger.New(cfg.Logger)
}
