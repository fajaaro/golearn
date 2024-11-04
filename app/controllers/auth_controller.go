package controllers

import (
	"fmt"
	"learn/app/helpers"
	"learn/app/models"
	"learn/config"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(data jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	tokenStr, _ := token.SignedString(SECRET)

	return tokenStr
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

	go func(email string) {
		appName := os.Getenv("APP_NAME")
		to := []string{email}
		subject := appName + " Registration"
		body := "You have successfully registered to " + appName + " application."
		_ = helpers.SendEmail(to, subject, body)
	}(entity.Email)

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
	tokenData := jwt.MapClaims{
		"token_type": "access_token",
		"sub": user.ID,
		"iat": time.Now().Unix(),
		"exp": accessTokenExpirationTime.Unix(),
	}

	accessToken := GenerateToken(tokenData)
	
	tokenData["token_type"] = "refresh_token"
	tokenData["exp"] = refreshTokenExpirationTime.Unix()
	refreshToken := GenerateToken(tokenData)

	resObj := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	res.Data = resObj
	c.JSON(200, res)
}

func RefreshToken(c *gin.Context) {
	res := models.Response{Success: true}
	req := map[string]string{"refresh_token": ""}
	c.ShouldBindJSON(&req)

	claims, err := DecodeToken(req["refresh_token"])
	if err != nil {
		errMsg := err.Error()
		res.Success = false
		res.Error = &errMsg
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	_, found := claims["token_type"]
	if !found || claims["token_type"] != "refresh_token" {
		msgErr := "Refresh Token is invalid"
		res.Success = false
		res.Error = &msgErr
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	accessTokenExpirationTime := time.Now().Add(10*time.Minute)
	tokenData := jwt.MapClaims{
		"token_type": "access_token",
		"sub": claims["sub"],
		"iat": time.Now().Unix(),
		"exp": accessTokenExpirationTime.Unix(),
	}
	accessToken := GenerateToken(tokenData)

	res.Data = gin.H{
		"access_token": accessToken,
	}

	c.JSON(http.StatusOK, res)
}
