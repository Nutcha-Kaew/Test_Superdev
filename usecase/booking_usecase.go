package usecase

import (
	"fiber-booking-system/models"
	"fiber-booking-system/repository"
	"fiber-booking-system/utils"
)

type BookingUseCase struct {
	Repository repository.BookingRepository
	Cache      utils.Cache
}

func NewBookingUseCase(repo repository.BookingRepository, cache utils.Cache) *BookingUseCase {
	return &BookingUseCase{Repository: repo, Cache: cache}
}

func (uc *BookingUseCase) CreateBooking(userID, serviceID int, price float64) (*models.Booking, error) {
	booking := &models.Booking{
		UserID:    userID,
		ServiceID: serviceID,
		Price:     price,
		Status:    "pending",
	}

	// Save booking to repository (mock or DB)
	uc.Repository.SaveBooking(booking)

	// Cache the booking
	uc.Cache.Save(booking)

	return booking, nil
}
