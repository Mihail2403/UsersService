package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	ctx.JSON(201, gin.H{"message": "User created successfully"})
}
