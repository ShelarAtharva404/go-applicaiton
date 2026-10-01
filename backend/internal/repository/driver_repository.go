package repository

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type DriverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (r *DriverRepository) GetAll() ([]models.Driver, error) {
	var drivers []models.Driver

	result := r.db.Find(&drivers)

	return drivers, result.Error
}

func (r *DriverRepository) GetByID(id uint) (*models.Driver, error) {
	var driver models.Driver

	result := r.db.First(&driver, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &driver, nil
}

func (r *DriverRepository) Create(driver *models.Driver) error {
	return r.db.Create(driver).Error
}

func (r *DriverRepository) Update(driver *models.Driver) error {
	return r.db.Save(driver).Error
}

func (r *DriverRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Driver{}, id)

	return result.Error
}
