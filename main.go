package main

import (
	"github.com/fahmi-saputra/golang-restapi-gin/controllers/productcontroller"
	"github.com/fahmi-saputra/golang-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	models.ConnecDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()


}