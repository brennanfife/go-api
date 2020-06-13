package main

import (
	"github.com/brennanfife/go-api/controllers"
	"github.com/brennanfife/go-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
