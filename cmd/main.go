package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"fiber-booking-system/docs"
	"fiber-booking-system/router"
	"fiber-booking-system/middleware"
	"fiber-booking-system/repository"
	"fiber-booking-system/usecase"
	"fiber-booking-system/handler"
)

func main() {
	app := fiber.New()

	// Initialize repositories, usecases, and handlers
	cache := utils.NewInMemoryCache()
	repo := repository.NewMockBookingRepo()
	usecase := usecase.NewBookingUseCase(repo, cache)
	handler := handler.NewBookingHandler(usecase)

	// Initialize router
	router.Setup(app, handler)

	// Swagger setup
	docs.SwaggerInfo.Host = "localhost:3000"
	app.Get("/swagger/*", swagger.Handler)

	// Start server
	app.Listen(":3000")
}
