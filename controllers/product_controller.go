package controllers

import (
	"learn/helpers"
	"learn/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    helpers.Create[models.Product](c)
}

func GetProducts(c *gin.Context) {
    helpers.GetAll[models.Product](c)
}

func GetProductByID(c *gin.Context) {
    helpers.GetByID[models.Product](c)
}

func UpdateProduct(c *gin.Context) {
    helpers.Update[models.Product](c)
}

func DeleteProduct(c *gin.Context) {
    helpers.Delete[models.Product](c)
}

func RestoreProduct(c *gin.Context) {
    helpers.Restore[models.Product](c)
}
