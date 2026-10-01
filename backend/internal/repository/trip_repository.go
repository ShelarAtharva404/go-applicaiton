package repository

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) GetAll() ([]models.Trip, error) {
	var trips []models.Trip

	result := r.db.Find(&trips)

	return trips, result.Error
}

func (r *TripRepository) GetByDriverID(driverID uint) ([]models.Trip, error) {
	var trips []models.Trip

	result := r.db.
		Where("driver_id = ?", driverID).
		Find(&trips)

	return trips, result.Error
}

func (r *TripRepository) GetByID(id uint) (*models.Trip, error) {
	var trip models.Trip

	result := r.db.First(&trip, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &trip, nil
}

func (r *TripRepository) Create(trip *models.Trip) error {
	return r.db.Create(trip).Error
}

func (r *TripRepository) Update(trip *models.Trip) error {
	return r.db.Save(trip).Error
}
