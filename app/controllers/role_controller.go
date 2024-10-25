package controllers

import (
	"learn/app/helpers"
	"learn/app/models"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
    helpers.Create[models.Role](c)
}

func GetRoles(c *gin.Context) {
    helpers.GetAll[models.Role](c)
}

func GetRoleByID(c *gin.Context) {
    helpers.GetByID[models.Role](c)
}

func UpdateRole(c *gin.Context) {
    helpers.Update[models.Role](c)
}

func DeleteRole(c *gin.Context) {
    helpers.Delete[models.Role](c)
}

func RestoreRole(c *gin.Context) {
    helpers.Restore[models.Role](c)
}
