package touriscontroller

import (
	"backend/middleware"
	tourisService "backend/service/tourisService"
	tourisDto "backend/dto/tourisDto"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

type TourisController interface{}
type tourisController struct {
	tourisServ tourisService.TourisService
}

func NewAdminController(tourisService tourisService.TourisService) *tourisController {
	return &tourisController{
		tourisServ: tourisService,
	}
}

func (u *tourisController) GetProfileTouris(c echo.Context) error {

	userId, _ := middleware.ClaimData(c, "userID")
	conv_userId := userId.(float64)
	conv := uint(conv_userId)

	profile, err := u.tourisServ.GetProfileTouris(conv)
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

func (u *tourisController) GetAllReviewTouris(c echo.Context) error {
	userId, _ := middleware.ClaimData(c, "userID")
	conv_userId := userId.(float64)
	conv := uint(conv_userId)

	reviews, err := u.tourisServ.GetAllReviewTouris(conv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    reviews,
	})
}

func (u *tourisController) CreateReviewTouris(c echo.Context) error {
	userId, _ := middleware.ClaimData(c, "userID")
	conv_userId := userId.(float64)
	conv := uint(conv_userId)
	var payloads tourisDto.CreateReviewDTO
	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "Invalid input",
			Code:    http.StatusBadRequest,
		})
	}
	payloads.UserId = conv
	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
	if err := u.tourisServ.CreateReviewTouris(payloads); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}
	return c.JSON(http.StatusOK, utils.Response{
		Message: "Review created successfully",
		Code:    http.StatusOK,
	})

}

func (u *tourisController) UpdateReviewTouris(c echo.Context) error {
	id := c.FormValue("review_id")
	reviewId := utils.ParseUint(id)
	if reviewId != 0 {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "Invalid review_id",
			Code:    http.StatusBadRequest,
		})
	}

	if err := u.tourisServ.UpdateReviewTouris(reviewId); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Review updated successfully",
		Code:    http.StatusOK,
	})
}
func (u *tourisController) DeleteReviewTouris(c echo.Context) error {
	id := c.FormValue("review_id")
	reviewId := utils.ParseUint(id)
	if reviewId != 0 {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "Invalid review_id",
			Code:    http.StatusBadRequest,
		})
	}

	if err := u.tourisServ.DeleteReviewTouris(reviewId); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Review deleted successfully",
		Code:    http.StatusOK,
	})
}
