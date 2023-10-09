package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
}

func Handler(r *gin.RouterGroup, serv Service) {
	res := resource{service: serv}

	r.POST("/register", res.Register)
}

func (r resource) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.service.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated", "data": res})
}
