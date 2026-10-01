package models

import "time"

type Trip struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	VehicleID     uint       `json:"vehicle_id" gorm:"not null"`
	DriverID      uint       `json:"driver_id" gorm:"not null"`
	StartLocation string     `json:"start_location" gorm:"not null"`
	EndLocation   string     `json:"end_location" gorm:"not null"`
	Status        string     `json:"status" gorm:"not null;default:'PLANNED'"`
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
