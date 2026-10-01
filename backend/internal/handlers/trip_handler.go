package handlers

import (
	"net/http"
	"strconv"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TripHandler struct {
	service *services.TripService
	db      *gorm.DB
}

func NewTripHandler(service *services.TripService, db *gorm.DB) *TripHandler {
	return &TripHandler{
		service: service,
		db:      db,
	}
}

func (h *TripHandler) GetTrips(c *gin.Context) {
	roleValue, exists := c.Get("role")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		return
	}

	role := roleValue.(string)

	var trips []models.Trip
	var err error

	if role == "ADMIN" || role == "MANAGER" {
		// Admin and Manager can see all trips.
		trips, err = h.service.GetAllTrips()
	} else {
		// Driver can see only trips assigned to their driver profile.
		userIDValue, exists := c.Get("user_id")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user identity not found",
			})
			return
		}

		userID := userIDValue.(uint)

		// Find the driver profile linked to the logged-in user.
		var driver models.Driver

		result := h.db.
			Where("user_id = ?", userID).
			First(&driver)

		if result.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "driver profile not found",
			})
			return
		}

		trips, err = h.service.GetTripsByDriverID(driver.ID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch trips",
		})
		return
	}

	c.JSON(http.StatusOK, trips)
}

func (h *TripHandler) GetTrip(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid trip ID",
		})
		return
	}

	trip, err := h.service.GetTripByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "trip not found",
		})
		return
	}

	c.JSON(http.StatusOK, trip)
}

func (h *TripHandler) CreateTrip(c *gin.Context) {
	var trip models.Trip

	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if trip.VehicleID == 0 ||
		trip.DriverID == 0 ||
		trip.StartLocation == "" ||
		trip.EndLocation == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "vehicle_id, driver_id, start_location and end_location are required",
		})
		return
	}

	if trip.Status == "" {
		trip.Status = "PLANNED"
	}

	if err := h.service.CreateTrip(&trip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create trip",
		})
		return
	}

	c.JSON(http.StatusCreated, trip)
}

func (h *TripHandler) UpdateTrip(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid trip ID",
		})
		return
	}

	trip, err := h.service.GetTripByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "trip not found",
		})
		return
	}

	var request models.Trip

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	trip.VehicleID = request.VehicleID
	trip.DriverID = request.DriverID
	trip.StartLocation = request.StartLocation
	trip.EndLocation = request.EndLocation
	trip.Status = request.Status
	trip.StartedAt = request.StartedAt
	trip.CompletedAt = request.CompletedAt

	if err := h.service.UpdateTrip(trip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update trip",
		})
		return
	}

	c.JSON(http.StatusOK, trip)
}
