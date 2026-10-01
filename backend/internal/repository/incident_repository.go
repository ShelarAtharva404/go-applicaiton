package repository

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) GetAll() ([]models.Incident, error) {
	var incidents []models.Incident

	result := r.db.Find(&incidents)

	return incidents, result.Error
}

func (r *IncidentRepository) GetByID(id uint) (*models.Incident, error) {
	var incident models.Incident

	result := r.db.First(&incident, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &incident, nil
}

func (r *IncidentRepository) Create(incident *models.Incident) error {
	return r.db.Create(incident).Error
}

func (r *IncidentRepository) Update(incident *models.Incident) error {
	return r.db.Save(incident).Error
}
