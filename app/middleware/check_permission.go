package middleware

import (
	"learn/app/models"
	"learn/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPermission(actionType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := models.Response{Success: true}

		user, _ := c.Get("user")
		currentUser := user.(models.User)

		var permission models.Permission
		if err := config.DB.Model(&models.Permission{}).Where("name = ?", actionType).First(&permission).Error; err != nil {
			c.Next()
			return
		}

		isAllowed := false
	
		var userPermission models.UserPermission
		if err := config.DB.Model(&models.UserPermission{}).Where("user_id = ? and permission_id = ?", currentUser.ID, permission.ID).First(&userPermission).Error; err == nil {
			isAllowed = true
		}

		var rolesId []uint
		config.DB.Model(&models.UserRole{}).Where("user_id = ?", currentUser.ID).Pluck("role_id", &rolesId)

		var count int64
		config.DB.Model(&models.RolePermission{}).Where("role_id in ? and permission_id = ?", rolesId, permission.ID).Count(&count)

		if count > 0 {
			isAllowed = true
		}

		if (!isAllowed) {
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