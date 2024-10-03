package entrypoints

import "github.com/gin-gonic/gin"

// RegisterRoutes registers the routes for the entity
func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", healthCheckHandler)
}

// @Summary Health Check
// @Description Returns the health status of the API
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "healthy"})
}
