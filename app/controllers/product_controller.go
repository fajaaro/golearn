package controllers

import (
	"learn/app/helpers"
	"learn/app/models"
	"learn/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    res := models.Response{Success: true}

	var entity models.Product
	if err := c.Bind(&entity); err != nil {
		err := "Error parsing data: " + err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}	

    form, _ := c.MultipartForm()
    images := form.File["image"]
    if len(images) > 0 {
        image := images[0]
        filepath, err := helpers.UploadFile(c, image, "products")
        if err == nil {
            entity.ImagePath = filepath
        }
    }

	if err := config.DB.Create(&entity).Error; err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = entity
	c.JSON(http.StatusOK, res)
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
