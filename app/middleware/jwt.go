package middleware

import (
	"learn/app/controllers"
	"learn/app/models"
	"learn/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := models.Response{Success: true}

		if len(strings.Split(c.Request.Header.Get("Authorization"), " ")) != 2 {
			errorMsg := "Invalid token"
			res.Success = false
			res.Error = &errorMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		accessToken := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]

		claims, err := controllers.DecodeToken(accessToken)
		if err != nil {
			errMsg := err.Error()
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if claims["token_type"] != "access_token" {
			errMsg := "Invalid token"
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		var user models.User
		if err := config.DB.Where("id = ?", claims["sub"]).First(&user).Error; err != nil {
			errMsg := "Invalid token user " + claims["sub"].(string)
			res.Success = false
			res.Error = &errMsg
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		c.Set("user", user)
	
		c.Next()
	}
}