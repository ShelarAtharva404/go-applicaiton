package models

import "time"

type Incident struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	VehicleID   uint       `json:"vehicle_id" gorm:"not null"`
	DriverID    *uint      `json:"driver_id"`
	Type        string     `json:"type" gorm:"not null"`
	Description string     `json:"description"`
	Severity    string     `json:"severity" gorm:"not null;default:'LOW'"`
	Status      string     `json:"status" gorm:"not null;default:'OPEN'"`
	OccurredAt  time.Time  `json:"occurred_at"`
	ResolvedAt  *time.Time `json:"resolved_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
