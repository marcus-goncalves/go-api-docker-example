package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupportRoutes struct {
}

// @Router /support/ping [get]
// @Tags Support
// @Description Check if API is up returning "OK"
func (s *SupportRoutes) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
