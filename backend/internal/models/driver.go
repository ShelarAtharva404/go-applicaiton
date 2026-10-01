package models

import "time"

type Driver struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        *uint     `json:"user_id" gorm:"unique"`
	Name          string    `json:"name" gorm:"not null"`
	Email         string    `json:"email" gorm:"unique;not null"`
	Phone         string    `json:"phone" gorm:"unique;not null"`
	LicenseNumber string    `json:"license_number" gorm:"unique;not null"`
	Status        string    `json:"status" gorm:"not null;default:'ACTIVE'"`
	VehicleID     *uint     `json:"vehicle_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
