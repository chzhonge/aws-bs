package check

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (h Handler) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "#%s", "health check")
}
