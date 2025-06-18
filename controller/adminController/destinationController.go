package adminController

import (
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (u *adminController) GetAllDestination(c echo.Context) error {
	res, err := u.adminServ.GetAllDestination()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
