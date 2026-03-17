package user

import (
	"net/http"
	"strconv"
	"user-flow/internal/app/core/helpers/errorhandler"
	response "user-flow/internal/app/core/helpers/response"
	dto "user-flow/internal/app/domain/core/dto/requests"
	filter "user-flow/internal/app/domain/core/filter/user"
	service "user-flow/internal/app/domain/services/user"

	"github.com/gin-gonic/gin"
)

var errStatusMap = map[error]int{
	service.ErrNotFound: http.StatusNotFound,
}

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
		errorhandler.AbortWithError(func(code int, obj any) { ctx.JSON(code, obj) }, err, errStatusMap)
		errorhandler.FailOnError(err, "Show user service error")

		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) Get(ctx *gin.Context) {
	var getDTO dto.GetDTO

	if err := ctx.ShouldBind(&getDTO); err != nil {
		errorhandler.FailOnError(err, "Validation error")

		return
	}

	users, err := h.service.Get(filter.GetFilter{
		Keyword: getDTO.Keyword,
		Page:    getDTO.Page,
		PerPage: getDTO.PerPage,
	})
	if err != nil {
		errorhandler.AbortWithError(func(code int, obj any) { ctx.JSON(code, obj) }, err, errStatusMap)
		errorhandler.FailOnError(err, "Get user service error")

		return
	}

	ctx.JSON(http.StatusOK, users)
}
