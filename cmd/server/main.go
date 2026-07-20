package main

import (

	"github.com/Vophat289/go-commerce-system/internal/router"
	// "github.com/gin-gonic/gin"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  r := router.NewRouter()

  r.Run(":8080")
}