package services

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
)

type TripService struct {
	repository *repository.TripRepository
}

func NewTripService(repository *repository.TripRepository) *TripService {
	return &TripService{
		repository: repository,
	}
}

func (s *TripService) GetAllTrips() ([]models.Trip, error) {
	return s.repository.GetAll()
}

func (s *TripService) GetTripsByDriverID(driverID uint) ([]models.Trip, error) {
	return s.repository.GetByDriverID(driverID)
}

func (s *TripService) GetTripByID(id uint) (*models.Trip, error) {
	return s.repository.GetByID(id)
}

func (s *TripService) CreateTrip(trip *models.Trip) error {
	return s.repository.Create(trip)
}

func (s *TripService) UpdateTrip(trip *models.Trip) error {
	return s.repository.Update(trip)
}
