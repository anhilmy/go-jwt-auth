package main

import (
	"go-jwt-auth/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handler(r)

	r.Run()

}

func handler(router *gin.Engine) {
	api := router.Group("/api")

	v1 := api.Group("/v1")

	route.AuthHandler(v1)

}
