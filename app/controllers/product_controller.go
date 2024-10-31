package controllers

import (
	"fmt"
	"learn/app/helpers"
	"learn/app/models"
	"learn/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ImportExcelProduct(c *gin.Context) {
    res := models.Response{Success: true}

	file, err := c.FormFile("file")
	if err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	rows, err := helpers.ReadExcel(file)
	if err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	var rowsData [][]string
	if len(rows) > config.Constant.UploadExcelStartFromIndex {
		rowsData = rows[config.Constant.UploadExcelStartFromIndex:]
	} else {
		rowsData = [][]string{}
	} 

	excelIndex := helpers.ExtractModelExcelColIndexes(models.Product{})

	var errors []string
	totalInserted := 0
	for index, row := range rowsData {
		entity := make(map[string]interface{})	
		for field, colExcelIndex := range excelIndex {
			entity[field] = row[colExcelIndex]
		}
		entity["CreatedAt"] = time.Now()
		entity["UpdatedAt"] = time.Now()

		if err := config.DB.Model(&models.Product{}).Create(entity).Error; err != nil {
			errors = append(errors, fmt.Sprintf("Row %d: %s", index + 4, err.Error()))
		} else {
			totalInserted++
		}
	}
	if len(errors) == 0 {
		errors = []string{}
	}

	res.Data = map[string]interface{}{
		"message": fmt.Sprintf("Successfully upload %d product data!", totalInserted),
		"errors": errors,
	}
	c.JSON(http.StatusOK, res)
}

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
        filepath, err := helpers.UploadFile(c, image, "private", "products")
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
