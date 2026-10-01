package repository

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type VehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{db: db}
}

func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle

	result := r.db.Find(&vehicles)

	return vehicles, result.Error
}

func (r *VehicleRepository) GetByID(id uint) (*models.Vehicle, error) {
	var vehicle models.Vehicle

	result := r.db.First(&vehicle, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &vehicle, nil
}

func (r *VehicleRepository) Create(vehicle *models.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *VehicleRepository) Update(vehicle *models.Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *VehicleRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Vehicle{}, id)

	return result.Error
}
