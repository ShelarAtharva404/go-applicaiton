package services

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
)

type DriverService struct {
	repository *repository.DriverRepository
}

func NewDriverService(repository *repository.DriverRepository) *DriverService {
	return &DriverService{
		repository: repository,
	}
}

func (s *DriverService) GetAllDrivers() ([]models.Driver, error) {
	return s.repository.GetAll()
}

func (s *DriverService) GetDriverByID(id uint) (*models.Driver, error) {
	return s.repository.GetByID(id)
}

func (s *DriverService) CreateDriver(driver *models.Driver) error {
	return s.repository.Create(driver)
}

func (s *DriverService) UpdateDriver(driver *models.Driver) error {
	return s.repository.Update(driver)
}

func (s *DriverService) DeleteDriver(id uint) error {
	return s.repository.Delete(id)
}
