package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DriverTripAccess(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			c.Abort()
			return
		}

		role := roleValue.(string)

		// ADMIN and MANAGER can access all trips.
		if role != "DRIVER" {
			c.Next()
			return
		}

		userIDValue, exists := c.Get("user_id")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user identity not found",
			})
			c.Abort()
			return
		}

		userID := userIDValue.(uint)

		tripID, err := strconv.ParseUint(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid trip ID",
			})
			c.Abort()
			return
		}

		var driver struct {
			ID uint
		}

		result := db.Table("drivers").
			Select("id").
			Where("user_id = ?", userID).
			First(&driver)

		if result.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "driver profile not found",
			})
			c.Abort()
			return
		}

		var trip struct {
			DriverID uint
		}

		result = db.Table("trips").
			Select("driver_id").
			Where("id = ?", tripID).
			First(&trip)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "trip not found",
			})
			c.Abort()
			return
		}

		if trip.DriverID != driver.ID {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you can only access your assigned trips",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
