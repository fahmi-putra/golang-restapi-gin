package productcontroller

import (
	"net/http"

	"github.com/fahmi-saputra/golang-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var product []models.Product

	models.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return

		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
	
}
func Create(c *gin.Context) {
	
}
func Update(c *gin.Context) {
	
}
func Delete(c *gin.Context) {
	
}