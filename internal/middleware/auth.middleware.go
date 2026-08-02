package middleware

import (
	"github.com/Vophat289/go-commerce-system/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if token != "valid-token"{
			response.ErrorRes(c, response.ErrInvalid, "")
			c.Abort()

			return  
		}
		c.Next()
	}
}