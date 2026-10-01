package repository

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (r *MaintenanceRepository) GetAll() ([]models.Maintenance, error) {
	var maintenance []models.Maintenance

	result := r.db.Find(&maintenance)

	return maintenance, result.Error
}

func (r *MaintenanceRepository) GetByID(id uint) (*models.Maintenance, error) {
	var maintenance models.Maintenance

	result := r.db.First(&maintenance, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &maintenance, nil
}

func (r *MaintenanceRepository) Create(maintenance *models.Maintenance) error {
	return r.db.Create(maintenance).Error
}

func (r *MaintenanceRepository) Update(maintenance *models.Maintenance) error {
	return r.db.Save(maintenance).Error
}
