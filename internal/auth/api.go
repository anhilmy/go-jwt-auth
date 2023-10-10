package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
}

func Handler(r *gin.RouterGroup, serv Service) {
	res := resource{service: serv}

	r.POST("/register", res.Register)
	r.POST("/login", res.Login)
	r.GET("/profile", res.Profile)
}

func (r resource) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
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

func (r resource) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := r.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated", "data": res})

}

func (r resource) Profile(c *gin.Context) {
	bearer := r.extractToken(c)

	if bearer == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	}

	res, err := r.service.Profile(c.Request.Context(), bearer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})

}

func (r resource) extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearer := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearer, " ")) == 2 {
		return strings.Split(bearer, " ")[1]
	}
	return ""
}
