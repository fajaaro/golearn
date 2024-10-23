package controllers

import (
	"learn/config"
	"learn/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	res := models.Response{Success: true}

	var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
		err := err.Error()
		res.Error = &err
        c.JSON(http.StatusBadRequest, res)
        return
    }

	if err := config.DB.Create(&product).Error; err != nil {
		err := err.Error()
		res.Error = &err
        c.JSON(http.StatusInternalServerError, res)
        return
    }
	
	res.Data = product

    c.JSON(http.StatusOK, res)
}

func GetProducts(c *gin.Context) {
	res := models.Response{Success: true}

    var products []models.Product
    config.DB.Find(&products)

	res.Data = products

    c.JSON(http.StatusOK, res)
}

func GetProductByID(c *gin.Context) {
	res := models.Response{Success: true}

	var product models.Product
    id := c.Param("id")
    if err := config.DB.First(&product, id).Error; err != nil {
		err := "Product not found"
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }

	res.Data = product

    c.JSON(http.StatusOK, res)
}

func UpdateProduct(c *gin.Context) {
	res := models.Response{Success: true}

	var product models.Product
    id := c.Param("id")
    if err := config.DB.First(&product, id).Error; err != nil {
		err := "Product not found"
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
		err := err.Error()
		res.Error = &err
        c.JSON(http.StatusBadRequest, res)
        return
    }
    config.DB.Save(&product)

	res.Data = product

    c.JSON(http.StatusOK, res)
}

func DeleteProduct(c *gin.Context) {
	res := models.Response{Success: true}

	var product models.Product
    id := c.Param("id")

    if err := config.DB.Unscoped().First(&product, id).Error; err != nil {
		err := "Product not found"
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }

    var message string
    if product.DeletedAt.Valid {
        config.DB.Unscoped().Delete(&product, id)

        message = "Product deleted permanently"
    } else {
        config.DB.Delete(&product, id)

        message = "Product deleted successfully"
    }

	res.Data = map[string]interface{}{
		"message": message,
	}

    c.JSON(http.StatusOK, res)
}
