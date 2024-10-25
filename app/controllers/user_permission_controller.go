package controllers

import (
	"learn/app/helpers"
	"learn/app/models"

	"github.com/gin-gonic/gin"
)

func CreateUserPermission(c *gin.Context) {
    helpers.Create[models.UserPermission](c)
}

func GetUserPermissions(c *gin.Context) {
    helpers.GetAll[models.UserPermission](c)
}

func GetUserPermissionByID(c *gin.Context) {
    helpers.GetByID[models.UserPermission](c)
}

func UpdateUserPermission(c *gin.Context) {
    helpers.Update[models.UserPermission](c)
}

func DeleteUserPermission(c *gin.Context) {
    helpers.Delete[models.UserPermission](c)
}
