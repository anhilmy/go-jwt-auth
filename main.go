package main

import (
	"flag"
	"go-jwt-auth/database"
	"go-jwt-auth/internal/auth"
	"go-jwt-auth/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

var flagConfig = flag.String("config", "./.local.yml", "config file path")

func main() {
	r := gin.Default()

	config, err := config.Load(*flagConfig)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Configuration file is not found")
	}

	db, err := database.OpenDB(config)
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Cannot open database")
	}

	registerHandler(r, db)

	r.Run()

}

func registerHandler(router *gin.Engine, db *database.Database) {
	api := router.Group("/api")

	v1 := api.Group("/v1")

	// init repo
	userRepo := auth.NewRepository(db)

	// init serv
	authServ := auth.NewService(userRepo)

	auth.Handler(v1, authServ)

}
