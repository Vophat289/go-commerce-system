package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong( *gin.Context){
	name := c.DefaultQuery("name","messi")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message":"pong.zzzz.kkkk" + name,
		"uid": uid,
		"users": []string{"cr7", "m10", "bin"},
	})
}