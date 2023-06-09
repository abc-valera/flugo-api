package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/abc-valera/flugo-api/docs/swagger" // docs generated by Swag CLI
	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/handler"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/middleware"
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

//	@host						flugo-api.fly.dev
//	@BasePath					/
//	@query.collection.format	multi
func RunAPI(
	PORT string,
	services *service.Services,
	repos *repository.Repositories,
	usecases *application.Usecases,
) error {
	handlers := handler.NewHandlers(repos, services, usecases)

	// Custom error handler
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// TODO: write custom logger middleware
	// Logger middleware
	app.Use(middleware.NewLoggerMiddleware(services.Logger))

	routes(app, services, handlers)

	return app.Listen(PORT)
}
