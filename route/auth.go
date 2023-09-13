package route

import (
	"go-jwt-auth/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthHandler(r *gin.RouterGroup) {
	r.POST("/register", Register)
}

func Register(c *gin.Context) {
	var req auth.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated"})
}
