package repository

import "fiber-booking-system/models"

type BookingRepository interface {
	SaveBooking(booking *models.Booking)
	GetBookingByID(id int) (*models.Booking, error)
}

type MockBookingRepo struct {
	Bookings map[int]*models.Booking
}

func NewMockBookingRepo() *MockBookingRepo {
	return &MockBookingRepo{
		Bookings: make(map[int]*models.Booking),
	}
}

func (repo *MockBookingRepo) SaveBooking(booking *models.Booking) {
	repo.Bookings[booking.ID] = booking
}

func (repo *MockBookingRepo) GetBookingByID(id int) (*models.Booking, error) {
	booking, exists := repo.Bookings[id]
	if !exists {
		return nil, fmt.Errorf("Booking not found")
	}
	return booking, nil
}