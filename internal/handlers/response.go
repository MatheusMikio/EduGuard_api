package handlers

import (
	"net/http"

	"github.com/MatheusMikio/eduGuard_api/internal/domain/dtos/auth"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg interface{}) {
	ctx.JSON(code, gin.H{
		"message":   msg,
		"ErrorCode": code,
	})
}

func SendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	if code == http.StatusNoContent {
		ctx.Status(code)
		return
	}

	ctx.JSON(code, gin.H{
		"message": op,
		"data":    data,
	})
}

func SendAuthSuccess(
	ctx *gin.Context,
	code int, accessToken string,
	refreshToken string,
	expiresIn int64,
	user auth.AuthUser,
) {
	ctx.JSON(code, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    expiresIn,
		"user":          user,
	})
}
