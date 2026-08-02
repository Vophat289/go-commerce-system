package router

import (
	"fmt"
	"net/http"

	"github.com/Vophat289/go-commerce-system/internal/controller"
	"github.com/Vophat289/go-commerce-system/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context){
	fmt.Println("before --> AA")
	c.Next()
	fmt.Println("after --> AA")
	}
	
}

func BB() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("before --> BB")
		c.Next()
		fmt.Println("after --> BB")
	}
}


func CC(c *gin.Context){
	
		fmt.Println("before --> CC")
		c.Next()
		fmt.Println("after --> CC")
	
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthenMiddleware(),AA(),BB(), CC)

	v1 := r.Group("/v1/2026")
	{
		v1.GET("/ping", Pong) // /v1/2026/ping
		v1.GET("/user/1", controller.NewUserController().GetUsersByID)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("/v2/2026")
	{
		v2.GET("/ping", Pong) // /v2/2026/ping
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "bin")

	// c.ShouldBindJSON()

	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "pong.hhh.ping" + name,
		"uid":     uid,
		"users":   []string{"haaland", "messi", "foden"},
	})
}
