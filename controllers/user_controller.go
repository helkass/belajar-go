package controllers

import (
	"belajar-go/dto/users"
	"belajar-go/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req users.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.CreateUser(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "success",
		"message": "User created",
		"data":    result,
	})
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.service.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Users fetched",
		"data":    users,
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	err := c.service.DeleteUser(idParam)

	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "succes",
		"message": "User deleted",
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var req users.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.UpdateUser(idParam, req)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status":  "success",
		"message": "User updated",
		"data":    result,
	})
}
