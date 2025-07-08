package router

import (
	"backend/controller/adminController"
	tourisController "backend/controller/tourisController"
	m "backend/middleware"
	"backend/repository/adminRepository"
	tourisRepository "backend/repository/tourisRepository"
	"backend/service/adminService"
	tourisService "backend/service/tourisService"
	"backend/utils"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {

	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// TODO REPOSITORY
	adminRepository := adminRepository.NewAdminRepository(db)
	tourisRepository := tourisRepository.NewTourisRepository(db)

	// TODO SERVICE
	adminService := adminService.NewAdminService(adminRepository)
	tourisService := tourisService.NewTourisService(tourisRepository)

	// TODO CONTROLLER
	adminController := adminController.NewAdminController(adminService)
	touriscontroller := tourisController.NewAdminController(tourisService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ping")
	})

	// TODO ROLE ROUTE
	v1 := e.Group("/v1")
	v1.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	v1.Use(m.RoleJwt(utils.GetRoleInt("ADMIN")))

	v2 := e.Group("/v2")
	v2.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))
	v2.Use(m.RoleJwt(utils.GetRoleInt("TOURIS")))

	// TODO ADMIN ROUTE
	// TODO AUTH

	// TODO PROFILE. Admin
	v1_profile := v1.Group("/profile")
	v1_profile.GET("/", adminController.GetProfileAdmin)
	// TODO PROFILE Touris
	v2_profile := v2.Group("/profile")
	v2_profile.GET("/", touriscontroller.GetProfileTouris)



	// TODO CRUD Destination
	v1_dest := v1.Group("/destination")
	v1_dest.POST("/", adminController.CreateDestination)
	v1_dest.PUT("/", adminController.UpdateDestination)
	v1_dest.DELETE("/", adminController.DeleteDestination)
	// v2_dest := v2.Group("/destination")

	// TODO FOR ALL
	e.POST("/auth/register", adminController.Register)
	e.POST("/auth/login", adminController.LoginAdmin)
	e.GET("/destination", adminController.GetAllDestination)
	e.GET("/destination/id", adminController.GetDestinationById)


	// TODO FOR TOURIS
	v2.GET("/review", touriscontroller.GetAllReviewTouris)
	v2.POST("/review", touriscontroller.CreateReviewTouris)
	v2.PUT("/review", touriscontroller.UpdateReviewTouris)
	v2.DELETE("/review", touriscontroller.DeleteReviewTouris)

}
