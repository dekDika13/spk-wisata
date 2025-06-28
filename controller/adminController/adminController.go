package adminController

import (
	adminDto "backend/dto/adminDto"
	"backend/service/adminService"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

type AdminController interface{}
type adminController struct {
	adminServ adminService.AdminService
}

func NewAdminController(adminService adminService.AdminService) *adminController {
	return &adminController{
		adminServ: adminService,
	}
}

// TODO REGISTER ADMIN
func (u *adminController) Register(c echo.Context) error {
	var payloads adminDto.RegisterInsertDTO

	// Ambil manual semua field form
	payloads.Role = 2
	payloads.Username = c.FormValue("username")
	payloads.Password = c.FormValue("password")
	payloads.Email = c.FormValue("email")
	payloads.Phone = c.FormValue("phone")
	payloads.FullName = c.FormValue("full_name")
	payloads.Bod = c.FormValue("bod")
	payloads.Address = c.FormValue("address")

	// Ambil file
	photo, err := c.FormFile("photo")
	if err == nil {
		payloads.Photo = photo
	}

	// Validasi
	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	// Proses register
	if err := u.adminServ.RegisterAdmin(payloads); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
	})
}

// TODO LOGIN ADMIN
func (u *adminController) LoginAdmin(c echo.Context) error {
	var payloads adminDto.LoginDTO

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.LoginAdmin(payloads)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.Response{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "login success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
