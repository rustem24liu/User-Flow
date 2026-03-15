package user

import (
	"net/http"
	"strconv"
	"user-flow/internal/app/core/helpers/errorhandler"
	response "user-flow/internal/app/core/helpers/response"
	service "user-flow/internal/app/domain/services/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Show(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(response.ServerError))
		errorhandler.FailOnError(err, "Validation error")

		return
	}

	user, err := h.service.Show(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(response.ServerError))
		errorhandler.FailOnError(err, "Show user service error")

		return
	}

	ctx.JSON(http.StatusOK, user)
}
