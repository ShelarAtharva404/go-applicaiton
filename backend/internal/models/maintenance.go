package models

import "time"

type Maintenance struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	VehicleID   uint       `json:"vehicle_id" gorm:"not null"`
	Type        string     `json:"type" gorm:"not null"`
	Description string     `json:"description"`
	Cost        float64    `json:"cost"`
	Status      string     `json:"status" gorm:"not null;default:'PENDING'"`
	ScheduledAt *time.Time `json:"scheduled_at"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
