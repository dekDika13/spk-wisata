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
func (u *adminController) GetDestinationById(c echo.Context) error {
	id := c.FormValue("destination_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.GetDestinationById(utils.StringToInt(id))
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
