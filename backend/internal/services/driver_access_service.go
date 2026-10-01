package services

import (
	"errors"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"gorm.io/gorm"
)

type DriverAccessService struct {
	db *gorm.DB
}

func NewDriverAccessService(db *gorm.DB) *DriverAccessService {
	return &DriverAccessService{
		db: db,
	}
}

func (s *DriverAccessService) GetDriverByUserID(userID uint) (*models.Driver, error) {
	var driver models.Driver

	result := s.db.Where("user_id = ?", userID).First(&driver)

	if result.Error != nil {
		return nil, result.Error
	}

	return &driver, nil
}

func (s *DriverAccessService) CanAccessVehicle(
	userID uint,
	vehicleID uint,
) (bool, error) {

	driver, err := s.GetDriverByUserID(userID)

	if err != nil {
		return false, err
	}

	if driver.VehicleID == nil {
		return false, nil
	}

	return *driver.VehicleID == vehicleID, nil
}

func (s *DriverAccessService) CanAccessTrip(
	userID uint,
	trip *models.Trip,
) (bool, error) {

	driver, err := s.GetDriverByUserID(userID)

	if err != nil {
		return false, err
	}

	return driver.ID == trip.DriverID, nil
}

func (s *DriverAccessService) ValidateDriver(userID uint) error {
	_, err := s.GetDriverByUserID(userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("driver profile not found")
		}

		return err
	}

	return nil
}
