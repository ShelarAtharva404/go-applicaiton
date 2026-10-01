package services

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
)

type VehicleService struct {
	repository *repository.VehicleRepository
}

func NewVehicleService(repository *repository.VehicleRepository) *VehicleService {
	return &VehicleService{
		repository: repository,
	}
}

func (s *VehicleService) GetAllVehicles() ([]models.Vehicle, error) {
	return s.repository.GetAll()
}

func (s *VehicleService) GetVehicleByID(id uint) (*models.Vehicle, error) {
	return s.repository.GetByID(id)
}

func (s *VehicleService) CreateVehicle(vehicle *models.Vehicle) error {
	return s.repository.Create(vehicle)
}

func (s *VehicleService) UpdateVehicle(vehicle *models.Vehicle) error {
	return s.repository.Update(vehicle)
}

func (s *VehicleService) DeleteVehicle(id uint) error {
	return s.repository.Delete(id)
}
