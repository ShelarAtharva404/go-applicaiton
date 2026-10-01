package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DashboardHandler struct {
	db *gorm.DB
}

func NewDashboardHandler(db *gorm.DB) *DashboardHandler {
	return &DashboardHandler{
		db: db,
	}
}

func (h *DashboardHandler) GetSummary(c *gin.Context) {
	var totalVehicles int64
	var activeVehicles int64
	var totalDrivers int64
	var activeDrivers int64
	var totalTrips int64
	var plannedTrips int64
	var completedTrips int64
	var openIncidents int64
	var pendingMaintenance int64

	if err := h.db.Table("vehicles").Count(&totalVehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch vehicle statistics",
		})
		return
	}

	if err := h.db.Table("vehicles").
		Where("status = ?", "ACTIVE").
		Count(&activeVehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch active vehicle statistics",
		})
		return
	}

	if err := h.db.Table("drivers").Count(&totalDrivers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch driver statistics",
		})
		return
	}

	if err := h.db.Table("drivers").
		Where("status = ?", "ACTIVE").
		Count(&activeDrivers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch active driver statistics",
		})
		return
	}

	if err := h.db.Table("trips").Count(&totalTrips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch trip statistics",
		})
		return
	}

	if err := h.db.Table("trips").
		Where("status = ?", "PLANNED").
		Count(&plannedTrips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch planned trip statistics",
		})
		return
	}

	if err := h.db.Table("trips").
		Where("status = ?", "COMPLETED").
		Count(&completedTrips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch completed trip statistics",
		})
		return
	}

	if err := h.db.Table("incidents").
		Where("status = ?", "OPEN").
		Count(&openIncidents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch incident statistics",
		})
		return
	}

	if err := h.db.Table("maintenances").
		Where("status = ?", "PENDING").
		Count(&pendingMaintenance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch maintenance statistics",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicles": gin.H{
			"total":  totalVehicles,
			"active": activeVehicles,
		},
		"drivers": gin.H{
			"total":  totalDrivers,
			"active": activeDrivers,
		},
		"trips": gin.H{
			"total":     totalTrips,
			"planned":   plannedTrips,
			"completed": completedTrips,
		},
		"incidents": gin.H{
			"open": openIncidents,
		},
		"maintenance": gin.H{
			"pending": pendingMaintenance,
		},
	})
}
