package helpers

import (
	"learn/config"
	"learn/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Create[T any](c *gin.Context) {
	res := models.Response{Success: true}

	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		err := err.Error()
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := config.DB.Create(&entity).Error; err != nil {
		err := err.Error()
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

	if err := q.Find(&entities).Error; err != nil {
		err := err.Error()
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
		res.Error = &err
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err := c.ShouldBindJSON(&entity); err != nil {
		err := err.Error()
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := config.DB.Save(&entity).Error; err != nil {
		err := err.Error()
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
    }

	res.Data = map[string]interface{}{
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
		res.Error = &err
        c.JSON(http.StatusNotFound, res)
        return
    }

	config.DB.Unscoped().Model(&entity).Update("deleted_at", nil)

	res.Data = map[string]interface{}{
		"message": "Entity restored successfully",
	}

	c.JSON(http.StatusOK, res)
}
