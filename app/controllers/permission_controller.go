package controllers

import (
	"learn/app/helpers"
	"learn/app/models"

	"github.com/gin-gonic/gin"
)

func CreatePermission(c *gin.Context) {
    helpers.Create[models.Permission](c)
}

func GetPermissions(c *gin.Context) {
    helpers.GetAll[models.Permission](c)
}

func GetPermissionByID(c *gin.Context) {
    helpers.GetByID[models.Permission](c)
}

func UpdatePermission(c *gin.Context) {
    helpers.Update[models.Permission](c)
}

func DeletePermission(c *gin.Context) {
    helpers.Delete[models.Permission](c)
}

func RestorePermission(c *gin.Context) {
    helpers.Restore[models.Permission](c)
}
