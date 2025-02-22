package router

import (
	"github.com/gofiber/fiber/v2"
	"fiber-booking-system/handler"
)

func Setup(app *fiber.App, handler *handler.BookingHandler) {
	app.Post("/bookings", handler.CreateBooking)
	app.Get("/bookings/:id", handler.GetBookingByID)
	app.Get("/bookings", handler.GetBookings)
	app.Delete("/bookings/:id", handler.CancelBooking)
}
