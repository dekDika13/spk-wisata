package adminController

import (
	adminDto "backend/dto/adminDto"
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

func (u *adminController) CreateDestination(c echo.Context) error {
	var payloads adminDto.DestinationInsertDTO

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}


	if err := u.adminServ.CreateDestination(payloads); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, utils.Response{
		Message: "destination created successfully",
		Code:    http.StatusCreated,
	})
}
func (u *adminController) UpdateDestination(c echo.Context) error {
	id := c.FormValue("destination_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
	}

	var payload adminDto.DestinationUpdateDTO
	payload.Name = c.FormValue("name")
	payload.Description = c.FormValue("description")
	payload.Image1 = c.FormValue("image1")
	payload.Image2 = c.FormValue("image2")
	payload.Image3 = c.FormValue("image3")
	payload.Image4 = c.FormValue("image4")
	payload.Price = utils.StringToInt(c.FormValue("price"))
	payload.Address = c.FormValue("address")
	payload.Location = c.FormValue("location")

	if err := u.adminServ.UpdateDestination(utils.StringToInt(id), payload); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "destination updated successfully",
		Code:    http.StatusOK,
	})
}
func (u *adminController) DeleteDestination(c echo.Context) error {
	id := c.FormValue("destination_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
	}

	if err := u.adminServ.DeleteDestination(utils.StringToInt(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "destination deleted successfully",
		Code:    http.StatusOK,
	})
}
