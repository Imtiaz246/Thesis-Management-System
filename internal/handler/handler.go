package handler

import (
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func GetUserUniIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*token.MyCustomClaims).UniversityId
}
