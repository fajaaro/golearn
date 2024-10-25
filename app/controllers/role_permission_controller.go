package controllers

import (
	"learn/app/helpers"
	"learn/app/models"

	"github.com/gin-gonic/gin"
)

func CreateRolePermission(c *gin.Context) {
    helpers.Create[models.RolePermission](c)
}

func GetRolePermissions(c *gin.Context) {
    helpers.GetAll[models.RolePermission](c)
}

func GetRolePermissionByID(c *gin.Context) {
    helpers.GetByID[models.RolePermission](c)
}

func UpdateRolePermission(c *gin.Context) {
    helpers.Update[models.RolePermission](c)
}

func DeleteRolePermission(c *gin.Context) {
    helpers.Delete[models.RolePermission](c)
}
