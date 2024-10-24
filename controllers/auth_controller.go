package controllers

import (
	"fmt"
	"learn/config"
	"learn/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(data gin.H, exp time.Time) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token_type": data["token_type"],
		"exp":        exp.Unix(),
	})

	tokenStr, _ := token.SignedString(SECRET)

	return tokenStr
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {} else {
		return nil, err
	}
	return claims, nil
}

func Register(c *gin.Context) {
	res := models.Response{Success: true}

	req := map[string]string{"email": "", "password": ""}
	if err := c.ShouldBindJSON(&req); err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var entity models.User
	var existingUser models.User

	if err := config.DB.Where("email = ?", strings.ToLower(req["email"])).First(&existingUser).Error; err == nil {
		errorMessage := "Email already exists"
		res.Success = false
		res.Error = &errorMessage
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"]), bcrypt.DefaultCost)
	if err != nil {
		errorMessage := "Failed to hash password"
		res.Success = false
		res.Error = &errorMessage
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	entity.Email = strings.ToLower(req["email"])
	entity.Password = string(hashedPassword)

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

func Login(c *gin.Context) {
	res := models.Response{Success: true}

	req := map[string]string{"email": "", "password": ""}
	if err := c.ShouldBindJSON(&req); err != nil {
		err := err.Error()
		res.Success = false
		res.Error = &err
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req["email"]).First(&user).Error; err != nil {
		errorMessage := "User not found"
		res.Success = false
		res.Error = &errorMessage
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req["password"])); err != nil {
		errorMessage := "Incorrect password"
		res.Success = false
		res.Error = &errorMessage
		c.JSON(http.StatusBadRequest, res)
		return
	}

	accessTokenExpirationTime := time.Now().Add(10*time.Minute)
	refreshTokenExpirationTime := time.Now().AddDate(0, 0, 5)
	tokenData := gin.H{
		"token_type": "access_token",
		"sub": user.ID,
		"iat": time.Now().String(),
		"exp": accessTokenExpirationTime.String(),
	}

	accessToken := GenerateToken(tokenData, accessTokenExpirationTime)
	
	tokenData["token_type"] = "refresh_token"
	tokenData["exp"] = refreshTokenExpirationTime.String()
	refreshToken := GenerateToken(tokenData, refreshTokenExpirationTime)

	resObj := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	res.Data = resObj
	c.JSON(200, res)
}
