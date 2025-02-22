package handler

import (
	"github.com/gofiber/fiber/v2"
	"fiber-booking-system/usecase"
	"fiber-booking-system/dto"
	"strconv"
)

type BookingHandler struct {
	UseCase *usecase.BookingUseCase
}

func NewBookingHandler(usecase *usecase.BookingUseCase) *BookingHandler {
	return &BookingHandler{UseCase: usecase}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var bookingDTO dto.BookingDTO
	if err := c.BodyParser(&bookingDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	booking, err := h.UseCase.CreateBooking(bookingDTO.UserID, bookingDTO.ServiceID, bookingDTO.Price)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.Status(201).JSON(booking)
}

func (h *BookingHandler) GetBookingByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	booking, err := h.UseCase.GetBookingByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Booking not found"})
	}

	return c.Status(200).JSON(booking)
}
