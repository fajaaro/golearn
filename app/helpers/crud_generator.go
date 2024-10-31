package helpers

import (
	"errors"
	"fmt"
	"learn/app/models"
	"learn/config"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func Create[T any](c *gin.Context) {
	res := models.Response{Success: true}

	contentType := Explode(";", c.Request.Header.Get("Content-Type"))[0]

	var entity T
	var err error

	if contentType == "application/json" {
		err = c.ShouldBindJSON(&entity);
	} else if contentType == "multipart/form-data" {
		err = c.Bind(&entity)
	} else {
		err = errors.New("unsupported content type")
	}

	if err != nil {
		err := "Error parsing data: " + err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
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

func GetAll[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entities []T
	q := config.DB.Unscoped()

	isDeleted := c.Query("is_deleted")
	if isDeleted == "1" {
		q = q.Where("deleted_at IS NOT NULL")
	} else if isDeleted == "0" {
		q = q.Where("deleted_at IS NULL")
	}

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "10")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	perPageInt, err := strconv.Atoi(perPage)
	if err != nil || perPageInt < 1 {
		perPageInt = 10
	}
	offset := (pageInt - 1) * perPageInt
	q = q.Offset(offset).Limit(perPageInt)

	orderBy := c.DefaultQuery("order_by", "id")
	orderType := c.DefaultQuery("order_type", "asc")
	orderType = strings.ToLower(orderType)
	if orderType != "asc" && orderType != "desc" {
		orderType = "asc"
	}
	q = q.Order(fmt.Sprintf("%s %s", orderBy, orderType))

	if err := q.Find(&entities).Error; err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = entities
	c.JSON(http.StatusOK, res)
}

func GetByID[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entity T
	id := c.Param("id")
	if err := config.DB.First(&entity, id).Error; err != nil {
		err := "Entity not found"
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusNotFound, res)
		return
	}

	res.Data = entity
	c.JSON(http.StatusOK, res)
}

func Update[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entity T
	id := c.Param("id")

	if err := config.DB.First(&entity, id).Error; err != nil {
		err := "Entity not found"
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := c.ShouldBindJSON(&entity); err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := config.DB.Save(&entity).Error; err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Data = entity
	c.JSON(http.StatusOK, res)
}

func Delete[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entity T
	id := c.Param("id")
	deleteType := c.Query("type")

    if err := config.DB.Unscoped().First(&entity, id).Error; err != nil {
		err := "Entity not found"
		res.Success = false
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }

    var message string
    if deleteType == "permanent" {
        config.DB.Unscoped().Delete(&entity, id)

        message = "Entity deleted permanently"
    } else if deleteType == "soft delete" {
        config.DB.Delete(&entity, id)

        message = "Entity deleted successfully"
    } else {
		err := "You must define delete type (param: type, options: permanent / soft delete)"
		res.Success = false
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
	}

	res.Data = gin.H{
		"message": message,
	}

	c.JSON(http.StatusOK, res)
}

func Restore[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entity T
	id := c.Param("id")

    if err := config.DB.Unscoped().First(&entity, id).Error; err != nil {
		err := "Entity not found"
		res.Success = false
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }

	config.DB.Unscoped().Model(&entity).Update("deleted_at", nil)

	res.Data = gin.H{
		"message": "Entity restored successfully",
	}

	c.JSON(http.StatusOK, res)
}
