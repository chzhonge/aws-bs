package main

import (
	"aws-bs/api/server/check"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	checkHandler := check.Handler{}

	r.GET("/health/check", checkHandler.HealthCheck)

	r.Run(":8080")
}
