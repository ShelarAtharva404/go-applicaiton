package services

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
)

type MaintenanceService struct {
	repository *repository.MaintenanceRepository
}

func NewMaintenanceService(repository *repository.MaintenanceRepository) *MaintenanceService {
	return &MaintenanceService{
		repository: repository,
	}
}

func (s *MaintenanceService) GetAllMaintenance() ([]models.Maintenance, error) {
	return s.repository.GetAll()
}

func (s *MaintenanceService) GetMaintenanceByID(id uint) (*models.Maintenance, error) {
	return s.repository.GetByID(id)
}

func (s *MaintenanceService) CreateMaintenance(maintenance *models.Maintenance) error {
	return s.repository.Create(maintenance)
}

func (s *MaintenanceService) UpdateMaintenance(maintenance *models.Maintenance) error {
	return s.repository.Update(maintenance)
}
