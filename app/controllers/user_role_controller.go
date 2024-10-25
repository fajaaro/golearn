package controllers

import (
	"learn/app/helpers"
	"learn/app/models"

	"github.com/gin-gonic/gin"
)

func CreateUserRole(c *gin.Context) {
    helpers.Create[models.UserRole](c)
}

func GetUserRoles(c *gin.Context) {
    helpers.GetAll[models.UserRole](c)
}

func GetUserRoleByID(c *gin.Context) {
    helpers.GetByID[models.UserRole](c)
}

func UpdateUserRole(c *gin.Context) {
    helpers.Update[models.UserRole](c)
}

func DeleteUserRole(c *gin.Context) {
    helpers.Delete[models.UserRole](c)
}
