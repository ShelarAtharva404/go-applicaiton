package handlers

import (
	"net/http"
	"strconv"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type MaintenanceHandler struct {
	service *services.MaintenanceService
}

func NewMaintenanceHandler(service *services.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{
		service: service,
	}
}

func (h *MaintenanceHandler) GetMaintenance(c *gin.Context) {
	maintenance, err := h.service.GetAllMaintenance()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch maintenance records",
		})
		return
	}

	c.JSON(http.StatusOK, maintenance)
}

func (h *MaintenanceHandler) GetMaintenanceByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid maintenance ID",
		})
		return
	}

	maintenance, err := h.service.GetMaintenanceByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "maintenance record not found",
		})
		return
	}

	c.JSON(http.StatusOK, maintenance)
}

func (h *MaintenanceHandler) CreateMaintenance(c *gin.Context) {
	var maintenance models.Maintenance

	if err := c.ShouldBindJSON(&maintenance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if maintenance.VehicleID == 0 || maintenance.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "vehicle_id and type are required",
		})
		return
	}

	if maintenance.Status == "" {
		maintenance.Status = "PENDING"
	}

	if err := h.service.CreateMaintenance(&maintenance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create maintenance record",
		})
		return
	}

	c.JSON(http.StatusCreated, maintenance)
}

func (h *MaintenanceHandler) UpdateMaintenance(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid maintenance ID",
		})
		return
	}

	maintenance, err := h.service.GetMaintenanceByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "maintenance record not found",
		})
		return
	}

	var request models.Maintenance

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	maintenance.VehicleID = request.VehicleID
	maintenance.Type = request.Type
	maintenance.Description = request.Description
	maintenance.Cost = request.Cost
	maintenance.Status = request.Status
	maintenance.ScheduledAt = request.ScheduledAt
	maintenance.CompletedAt = request.CompletedAt

	if err := h.service.UpdateMaintenance(maintenance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update maintenance record",
		})
		return
	}

	c.JSON(http.StatusOK, maintenance)
}
