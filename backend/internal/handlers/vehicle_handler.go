package handlers

import (
	"net/http"
	"strconv"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler(service *services.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		service: service,
	}
}

func (h *VehicleHandler) GetVehicles(c *gin.Context) {
	vehicles, err := h.service.GetAllVehicles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch vehicles",
		})
		return
	}

	c.JSON(http.StatusOK, vehicles)
}

func (h *VehicleHandler) GetVehicle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid vehicle ID",
		})
		return
	}

	vehicle, err := h.service.GetVehicleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vehicle not found",
		})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (h *VehicleHandler) CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if vehicle.RegistrationNumber == "" ||
		vehicle.Make == "" ||
		vehicle.Model == "" ||
		vehicle.VehicleType == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "registration_number, make, model and vehicle_type are required",
		})
		return
	}

	if err := h.service.CreateVehicle(&vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create vehicle",
		})
		return
	}

	c.JSON(http.StatusCreated, vehicle)
}

func (h *VehicleHandler) UpdateVehicle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid vehicle ID",
		})
		return
	}

	vehicle, err := h.service.GetVehicleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vehicle not found",
		})
		return
	}

	var request models.Vehicle

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	vehicle.RegistrationNumber = request.RegistrationNumber
	vehicle.Make = request.Make
	vehicle.Model = request.Model
	vehicle.VehicleType = request.VehicleType
	vehicle.Status = request.Status

	if err := h.service.UpdateVehicle(vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update vehicle",
		})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (h *VehicleHandler) DeleteVehicle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid vehicle ID",
		})
		return
	}

	_, err = h.service.GetVehicleByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vehicle not found",
		})
		return
	}

	if err := h.service.DeleteVehicle(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete vehicle",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "vehicle deleted successfully",
	})
}
