package main

import (
	"log"

	"github.com/atharva/fleetpulse-backend/internal/auth"
	"github.com/atharva/fleetpulse-backend/internal/database"
	"github.com/atharva/fleetpulse-backend/internal/handlers"
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"
	"github.com/atharva/fleetpulse-backend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// =========================================================
	// DATABASE
	// =========================================================

	db := database.Connect()

	// =========================================================
	// DATABASE MIGRATIONS
	// =========================================================

	err := db.AutoMigrate(
		&models.Vehicle{},
		&models.Driver{},
		&models.Trip{},
		&models.Maintenance{},
		&models.Incident{},
		&models.User{},
	)

	if err != nil {
		log.Fatal("Database migration failed:", err)
	}

	// =========================================================
	// VEHICLE
	// =========================================================

	vehicleRepository := repository.NewVehicleRepository(db)
	vehicleService := services.NewVehicleService(vehicleRepository)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	// =========================================================
	// DRIVER
	// =========================================================

	driverRepository := repository.NewDriverRepository(db)
	driverService := services.NewDriverService(driverRepository)
	driverHandler := handlers.NewDriverHandler(driverService)

	// =========================================================
	// TRIP
	// =========================================================

	tripRepository := repository.NewTripRepository(db)
	tripService := services.NewTripService(tripRepository)
	tripHandler := handlers.NewTripHandler(tripService, db)

	// =========================================================
	// MAINTENANCE
	// =========================================================

	maintenanceRepository := repository.NewMaintenanceRepository(db)
	maintenanceService := services.NewMaintenanceService(maintenanceRepository)
	maintenanceHandler := handlers.NewMaintenanceHandler(maintenanceService)

	// =========================================================
	// INCIDENT
	// =========================================================

	incidentRepository := repository.NewIncidentRepository(db)
	incidentService := services.NewIncidentService(incidentRepository)
	incidentHandler := handlers.NewIncidentHandler(incidentService)

	// =========================================================
	// AUTHENTICATION
	// =========================================================

	userRepository := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	// =========================================================
	// DASHBOARD
	// =========================================================

	dashboardHandler := handlers.NewDashboardHandler(db)

	// =========================================================
	// GIN ROUTER
	// =========================================================

	router := gin.Default()

	// =========================================================
	// CORS
	// =========================================================

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	// =========================================================
	// PUBLIC ROUTES
	// =========================================================

	// Health check does not require authentication.
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "fleetpulse-backend",
		})
	})

	// Authentication does not require JWT.
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	// =========================================================
	// PROTECTED API
	// =========================================================

	protected := api.Group("")
	protected.Use(auth.JWTMiddleware())

	// =========================================================
	// DASHBOARD
	// =========================================================

	// Dashboard now requires authentication.
	protected.GET("/dashboard/summary", dashboardHandler.GetSummary)

	// =========================================================
	// VEHICLES
	// =========================================================

	vehicles := protected.Group("/vehicles")

	// ADMIN / MANAGER / DRIVER can list vehicles.
	vehicles.GET("", vehicleHandler.GetVehicles)

	// Individual vehicle access.
	// DRIVER can only access assigned vehicle.
	vehicleRead := vehicles.Group("")
	vehicleRead.Use(auth.DriverVehicleAccess(db))

	vehicleRead.GET("/:id", vehicleHandler.GetVehicle)

	// Only ADMIN can create/update/delete vehicles.
	vehicleWrite := vehicles.Group("")
	vehicleWrite.Use(auth.RequireRole("ADMIN"))

	vehicleWrite.POST("", vehicleHandler.CreateVehicle)
	vehicleWrite.PUT("/:id", vehicleHandler.UpdateVehicle)
	vehicleWrite.DELETE("/:id", vehicleHandler.DeleteVehicle)

	// =========================================================
	// DRIVERS
	// =========================================================

	drivers := protected.Group("/drivers")

	// All authenticated users can view drivers.
	drivers.GET("", driverHandler.GetDrivers)
	drivers.GET("/:id", driverHandler.GetDriver)

	// Only ADMIN can manage drivers.
	driverWrite := drivers.Group("")
	driverWrite.Use(auth.RequireRole("ADMIN"))

	driverWrite.POST("", driverHandler.CreateDriver)
	driverWrite.PUT("/:id", driverHandler.UpdateDriver)
	driverWrite.DELETE("/:id", driverHandler.DeleteDriver)

	// =========================================================
	// TRIPS
	// =========================================================

	trips := protected.Group("/trips")

	// ADMIN and MANAGER see all trips.
	// DRIVER receives only their assigned trips
	// through TripHandler.GetTrips().
	trips.GET("", tripHandler.GetTrips)

	// Individual trip.
	// DRIVER can only access their own assigned trip.
	tripRead := trips.Group("")
	tripRead.Use(auth.DriverTripAccess(db))

	tripRead.GET("/:id", tripHandler.GetTrip)

	// ADMIN and MANAGER can create/update trips.
	tripWrite := trips.Group("")
	tripWrite.Use(auth.RequireRole("ADMIN", "MANAGER"))

	tripWrite.POST("", tripHandler.CreateTrip)
	tripWrite.PUT("/:id", tripHandler.UpdateTrip)

	// =========================================================
	// MAINTENANCE
	// =========================================================

	maintenance := protected.Group("/maintenance")

	// All authenticated users can view maintenance.
	maintenance.GET("", maintenanceHandler.GetMaintenance)
	maintenance.GET("/:id", maintenanceHandler.GetMaintenanceByID)

	// ADMIN and MANAGER can create/update maintenance.
	maintenanceWrite := maintenance.Group("")
	maintenanceWrite.Use(auth.RequireRole("ADMIN", "MANAGER"))

	maintenanceWrite.POST("", maintenanceHandler.CreateMaintenance)
	maintenanceWrite.PUT("/:id", maintenanceHandler.UpdateMaintenance)

	// =========================================================
	// INCIDENTS
	// =========================================================

	incidents := protected.Group("/incidents")

	// All authenticated users can view incidents.
	incidents.GET("", incidentHandler.GetIncidents)
	incidents.GET("/:id", incidentHandler.GetIncident)

	// ADMIN and MANAGER can create/update incidents.
	incidentWrite := incidents.Group("")
	incidentWrite.Use(auth.RequireRole("ADMIN", "MANAGER"))

	incidentWrite.POST("", incidentHandler.CreateIncident)
	incidentWrite.PUT("/:id", incidentHandler.UpdateIncident)

	// =========================================================
	// START SERVER
	// =========================================================

	log.Println("FleetPulse backend running on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
