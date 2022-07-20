package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{}

func (Response) ResponsFmt(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "msg": msg, "data": data})
}
func (r Response) Success(ctx *gin.Context, data gin.H, msg string) {
	r.ResponsFmt(ctx, http.StatusOK, 200, data, msg)
}
func (r Response) Error(ctx *gin.Context, data gin.H, msg string) {
	r.ResponsFmt(ctx, http.StatusBadRequest, 400, data, msg)
}
