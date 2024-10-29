package middleware

import (
	"learn/app/models"
	"learn/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSuperAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := models.Response{Success: true}

		user, _ := c.Get("user")
		currentUser := user.(models.User)

		var role models.Role
		if err := config.DB.Model(&models.Role{}).Where("name = 'Super Admin'").First(&role).Error; err != nil {
			c.Next()
			return
		}
	
		var count int64
		config.DB.Model(&models.UserRole{}).Where("user_id = ? and role_id = ?", currentUser.ID, role.ID).Count(&count)

		if (count == 0) {
			errMsg := "Forbidden"
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}