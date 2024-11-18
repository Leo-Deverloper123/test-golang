package routes

import (
	"hospital-middleware/handlers"
	"hospital-middleware/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	staffHandler := handlers.NewStaffHandler(db)
	patientHandler := handlers.NewPatientHandler(db)

	// Public routes
	staff := r.Group("/staff")
	{
		staff.POST("/create", staffHandler.CreateStaff)
		staff.POST("/login", staffHandler.Login)
	}

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/patient/search", patientHandler.SearchPatient)
		protected.GET("/patient/search/:id", patientHandler.SearchExternalPatient)
	}
}