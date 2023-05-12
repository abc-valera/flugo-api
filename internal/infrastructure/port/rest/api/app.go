package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "github.com/abc-valera/flugo-api/docs/swagger" // docs generated by Swag CLI
	"github.com/abc-valera/flugo-api/internal/application/service"
	"github.com/abc-valera/flugo-api/internal/application/usecase"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/infrastructure/port/rest/handler"
	_ "github.com/lib/pq"
)

//	@title			Flugo
//	@version		0.1
//	@description	API for Flugo social network

//	@contact.name	API Support
//	@contact.url	https://github.com/abc-valera
//	@contact.email	valeriy.tymofieiev@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/abc-valera

// @host						localhost:3000
// @BasePath					/
// @query.collection.format	multi
func RunAPI(PORT string, packages *service.Services, repos *repository.Repositories, services *usecase.Usecases) error {
	handlers := handler.NewHandlers(repos, services)

	// Custom error handler
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// TODO: write custom logger middleware
	// Logger middleware
	app.Use(logger.New())

	routes(app, packages, handlers)

	return app.Listen(PORT)
}
