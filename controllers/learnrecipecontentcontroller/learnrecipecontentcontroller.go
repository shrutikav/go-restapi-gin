package learnrecipecontentcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shrutikav/go-restapi-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var learnrecipecontent []models.Learn_recipe_content
	models.DB.Find(&learnrecipecontent)
	c.JSON(http.StatusOK, gin.H{"learn_recipe_content": learnrecipecontent})
}

func Show(c *gin.Context) {
	var recipe models.Learn_recipe_content
	id := c.Param("id")

	if err := models.DB.First(&recipe, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": recipe})
}
