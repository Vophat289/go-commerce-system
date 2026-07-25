package controller

import (
	"net/http"

	"github.com/Vophat289/go-commerce-system/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc -> user controller
// us -> user service
// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUsersByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"cr7", "m10", "bin"},
	})
}