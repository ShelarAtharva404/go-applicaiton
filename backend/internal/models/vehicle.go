package models

import "time"

type Vehicle struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	RegistrationNumber string    `json:"registration_number" gorm:"unique;not null"`
	Make               string    `json:"make" gorm:"not null"`
	Model              string    `json:"model" gorm:"not null"`
	VehicleType        string    `json:"vehicle_type" gorm:"not null"`
	Status             string    `json:"status" gorm:"not null;default:'ACTIVE'"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
