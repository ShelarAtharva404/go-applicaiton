package handlers

import (
	"net/http"
	"strconv"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	service *services.DriverService
}

func NewDriverHandler(service *services.DriverService) *DriverHandler {
	return &DriverHandler{
		service: service,
	}
}

func (h *DriverHandler) GetDrivers(c *gin.Context) {
	drivers, err := h.service.GetAllDrivers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch drivers",
		})
		return
	}

	c.JSON(http.StatusOK, drivers)
}

func (h *DriverHandler) GetDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid driver ID",
		})
		return
	}

	driver, err := h.service.GetDriverByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "driver not found",
		})
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var driver models.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if driver.Name == "" ||
		driver.Email == "" ||
		driver.Phone == "" ||
		driver.LicenseNumber == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name, email, phone and license_number are required",
		})
		return
	}

	if err := h.service.CreateDriver(&driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create driver",
		})
		return
	}

	c.JSON(http.StatusCreated, driver)
}

func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid driver ID",
		})
		return
	}

	driver, err := h.service.GetDriverByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "driver not found",
		})
		return
	}

	var request models.Driver

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if request.Name == "" ||
		request.Email == "" ||
		request.Phone == "" ||
		request.LicenseNumber == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name, email, phone and license_number are required",
		})
		return
	}

	driver.Name = request.Name
	driver.Email = request.Email
	driver.Phone = request.Phone
	driver.LicenseNumber = request.LicenseNumber
	driver.Status = request.Status
	driver.VehicleID = request.VehicleID

	if err := h.service.UpdateDriver(driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update driver",
		})
		return
	}

	c.JSON(http.StatusOK, driver)
}

func (h *DriverHandler) DeleteDriver(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid driver ID",
		})
		return
	}

	_, err = h.service.GetDriverByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "driver not found",
		})
		return
	}

	if err := h.service.DeleteDriver(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete driver",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "driver deleted successfully",
	})
}
