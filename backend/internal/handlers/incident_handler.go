package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type IncidentHandler struct {
	service *services.IncidentService
}

func NewIncidentHandler(service *services.IncidentService) *IncidentHandler {
	return &IncidentHandler{
		service: service,
	}
}

func (h *IncidentHandler) GetIncidents(c *gin.Context) {
	incidents, err := h.service.GetAllIncidents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch incidents",
		})
		return
	}

	c.JSON(http.StatusOK, incidents)
}

func (h *IncidentHandler) GetIncident(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid incident ID",
		})
		return
	}

	incident, err := h.service.GetIncidentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "incident not found",
		})
		return
	}

	c.JSON(http.StatusOK, incident)
}

func (h *IncidentHandler) CreateIncident(c *gin.Context) {
	var incident models.Incident

	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if incident.VehicleID == 0 || incident.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "vehicle_id and type are required",
		})
		return
	}

	if incident.Severity == "" {
		incident.Severity = "LOW"
	}

	if incident.Status == "" {
		incident.Status = "OPEN"
	}

	if incident.OccurredAt.IsZero() {
		incident.OccurredAt = time.Now()
	}

	if err := h.service.CreateIncident(&incident); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create incident",
		})
		return
	}

	c.JSON(http.StatusCreated, incident)
}

func (h *IncidentHandler) UpdateIncident(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid incident ID",
		})
		return
	}

	incident, err := h.service.GetIncidentByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "incident not found",
		})
		return
	}

	var request models.Incident

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	incident.VehicleID = request.VehicleID
	incident.DriverID = request.DriverID
	incident.Type = request.Type
	incident.Description = request.Description
	incident.Severity = request.Severity
	incident.Status = request.Status

	if !request.OccurredAt.IsZero() {
		incident.OccurredAt = request.OccurredAt
	}

	if request.ResolvedAt != nil {
		incident.ResolvedAt = request.ResolvedAt
	}
	if incident.Status == "RESOLVED" && incident.ResolvedAt == nil {
		now := time.Now()
		incident.ResolvedAt = &now
	}

	if err := h.service.UpdateIncident(incident); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update incident",
		})
		return
	}

	c.JSON(http.StatusOK, incident)
}
