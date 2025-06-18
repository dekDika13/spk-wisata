package router

import (
	"backend/controller/adminController"
	m "backend/middleware"
	"backend/repository/adminRepository"
	"backend/service/adminService"
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

	// TODO SERVICE
	adminService := adminService.NewAdminService(adminRepository)

	// TODO CONTROLLER
	adminController := adminController.NewAdminController(adminService)

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
	v1_reg := v1.Group("/reg")
	v1_reg.POST("/register", adminController.RegisterAdmin)
	e.POST("/auth/login", adminController.LoginAdmin)
	// TODO CRUD Destination
	v1_dest := v1.Group("/destination")
	v1_dest.GET("/all", adminController.GetAllDestination)
	// v1_dest.GET("/by-id/", adminController.GetDestinationById)
	// v1_dest.POST("/create", adminController.CreateDestination)
	// v1_dest.PUT("/update/", adminController.UpdateDestination)
	// v1_dest.DELETE("/delete/", adminController.DeleteDestination)

	//TODO CRUD CRITERIA
	// v1_crit := v1.Group("/criteria")
	// v1_crit.GET("/all", adminController.GetCriteriaAll)
	// v1_crit.GET("/by-id/", adminController.GetCriteriaById)
	// v1_crit.POST("/create", adminController.CreateCriteria)
	// v1_crit.PUT("/update/", adminController.UpdateCriteria)
	// v1_crit.DELETE("/delete/", adminController.DeleteCriteria)

	// TODO DETAIL Criteria
	// v1_crit_det := v1_crit.Group("/detail")
	// v1_crit_det.GET("/all", adminController.GetCriteriaDetailAll)
	// v1_crit_det.GET("/by-id/", adminController.GetCriteriaDetailById)

	// TODO DETAIL Review
	// v1_rev_det := v1.Group("/review/detail")
	// v1_rev_det.GET("/all", adminController.GetReviewDetailAll)
	// v1_rev_det.GET("/by-id/", adminController.GetReviewDetailById)
}
