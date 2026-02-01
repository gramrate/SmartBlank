package lobby

import (
	"backend/internal/domain/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req dto.CreateLobbyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := h.validator.ValidateCreateLobbyRequest(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//if err := h.validator.ValidateData(req); err != nil {
	//	return c.JSON(http.StatusBadRequest, dto.HTTPStatus{
	//		Code:    http.StatusBadRequest,
	//		Message: err.Error(),
	//	})
	//}

	resp, err := h.lobbyService.Create(c.Request().Context(), &req)
	switch {
	case err != nil:
		return c.JSON(http.StatusInternalServerError, dto.HTTPStatus{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, resp)

}
