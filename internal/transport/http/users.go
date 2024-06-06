package handler

import (
	"log"
	"users-service/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	var input entity.User
	err := ctx.BindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": "invalid body"})
		return
	}
	h.service.Users.Create(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "error creating user"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
}

func (h *Handler) GetUsersByManyId(ctx *gin.Context) {

}

func (h *Handler) GetById(ctx *gin.Context) {
}

func (h *Handler) UpdateUserById(ctx *gin.Context) {
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
}
