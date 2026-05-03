package base

import (
	"net/http"
	"strconv"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/models"
	"github.com/MatheusMikio/eduGuard_api/internal/handlers"
	"github.com/gin-gonic/gin"
)

const (
	MaxSize     uint64 = 100
	DefaultPage string = "1"
	DefaultSize string = "12"
)

func parsePagination(ctx *gin.Context) (uint64, uint64, bool) {
	pageStr := ctx.DefaultQuery("page", DefaultPage)
	sizeStr := ctx.DefaultQuery("size", DefaultSize)

	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil || page == 0 {
		message := "Parametro 'page' inválido, deve ser maior que 0"
		handlers.SendError(ctx, http.StatusBadRequest, message)
		return 0, 0, false
	}

	size, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil || size == 0 {
		message := "Parametro 'size' inválido, deve ser maior que 0"
		handlers.SendError(ctx, http.StatusBadRequest, message)
		return 0, 0, false
	}

	return page, size, true
}

func GetAllAdminHandler[TResponse any](
	getAllAdmin func(page uint64, size uint64) ([]*TResponse, *models.ErrorMessage),
	op string,
	errorPrefix string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, size, ok := parsePagination(ctx)
		if !ok {
			return
		}

		response, err := getAllAdmin(page, size)
		if err != nil {
			handlers.SendError(ctx, http.StatusInternalServerError, errorPrefix)
			return
		}
		handlers.SendSuccess(ctx, http.StatusOK, op, response)
	}
}

func GetByIdHandler[TResponse any](
	getById func(id uint) (*TResponse, *models.ErrorMessage),
	op string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil || id == 0 {
			message := "ID inválido, deve ser maior que 0."
			handlers.SendError(ctx, http.StatusBadRequest, message)
			return
		}

		entity, erro := getById(uint(id))
		if erro != nil {
			handlers.SendError(ctx, http.StatusNotFound, erro)
			return
		}
		handlers.SendSuccess(ctx, http.StatusOK, op, entity)
	}
}
