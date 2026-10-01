package services

import (
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
)

type IncidentService struct {
	repository *repository.IncidentRepository
}

func NewIncidentService(repository *repository.IncidentRepository) *IncidentService {
	return &IncidentService{
		repository: repository,
	}
}

func (s *IncidentService) GetAllIncidents() ([]models.Incident, error) {
	return s.repository.GetAll()
}

func (s *IncidentService) GetIncidentByID(id uint) (*models.Incident, error) {
	return s.repository.GetByID(id)
}

func (s *IncidentService) CreateIncident(incident *models.Incident) error {
	return s.repository.Create(incident)
}

func (s *IncidentService) UpdateIncident(incident *models.Incident) error {
	return s.repository.Update(incident)
}
