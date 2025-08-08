package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nostojic/gontacts/app"
	handlers "github.com/nostojic/gontacts/app/handlers/users"
	"github.com/nostojic/gontacts/config"
	"github.com/nostojic/gontacts/db"
)

func init() {
	config.LoadEnvs()
}

func main() {
	// connect to pg
	db, err := db.ConnectToDb()
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}
	defer db.Close()

	app := &app.App{Db: db}
	userHandler := &handlers.UserHandler{App: app}

	// start gin
	router := gin.Default()

	// user routes
	router.POST("/user", userHandler.UserCreate)
	router.DELETE("/user/:user_id", userHandler.UserDelete)

  router.Run()
}