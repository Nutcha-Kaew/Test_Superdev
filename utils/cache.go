package utils

import "fiber-booking-system/models"

type Cache interface {
	Save(booking *models.booking)
	Get(id int) (*models.Booking, bool)
}

type InMemoryCache struct {
	Bookings map[int]*models.Booking
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{Bookings: make(map[int]*models.Booking)}
}

func (c *InMemoryCache) Save(booking *models.Booking) {
	c.Bookings[booking.ID] = booking
}

func (c *InMemoryCache) Get(id int) (*models.Booking, bool) {
	booking, exists := c.Bookings[id]
	return booking, exists
}