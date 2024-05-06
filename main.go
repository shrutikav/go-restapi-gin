package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shrutikav/go-restapi-gin/controllers/learnrecipecontentcontroller"
	"github.com/shrutikav/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/learn_recipe_content", learnrecipecontentcontroller.Index)
	r.GET("/api/learn_recipe_content/:id", learnrecipecontentcontroller.Show)

	r.Run()
}
