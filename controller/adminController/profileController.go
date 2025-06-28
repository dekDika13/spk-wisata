package adminController

import (
	"backend/middleware"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (u *adminController) GetProfileAdmin(c echo.Context) error {

	userId, _ := middleware.ClaimData(c, "userID")
	conv_userId := userId.(float64)
	conv := uint(conv_userId)

	profile, err := u.adminServ.GetProfileAdmin(conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    profile,
	})
}
