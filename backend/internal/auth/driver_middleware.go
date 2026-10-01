package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DriverVehicleAccess(db *gorm.DB) gin.HandlerFunc {
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

		// ADMIN and MANAGER have normal fleet access.
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

		vehicleID, err := strconv.ParseUint(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid vehicle ID",
			})
			c.Abort()
			return
		}

		var driver struct {
			ID        uint
			VehicleID *uint
		}

		result := db.Table("drivers").
			Select("id, vehicle_id").
			Where("user_id = ?", userID).
			First(&driver)

		if result.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "driver profile not found",
			})
			c.Abort()
			return
		}

		if driver.VehicleID == nil ||
			*driver.VehicleID != uint(vehicleID) {

			c.JSON(http.StatusForbidden, gin.H{
				"error": "you can only access your assigned vehicle",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
