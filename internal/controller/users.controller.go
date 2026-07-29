package controller

import (
	"github.com/Vophat289/go-commerce-system/internal/service"
	"github.com/Vophat289/go-commerce-system/pkg/response"
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
    // Gọi service lấy data
    result := uc.userService.GetInfoUser()
    
    response.SuccessRes(c, 20001, result)
}