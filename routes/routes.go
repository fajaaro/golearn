package routes

import (
	"learn/controllers"
	"learn/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        authRoutes := api.Group("/auth")
        authRoutes.POST("/register", controllers.Register)
        authRoutes.POST("/login", controllers.Login)
        authRoutes.POST("/refresh-token", controllers.RefreshToken)

        productRoutes := api.Group("/products")
        productRoutes.Use(middleware.JWTMiddleware())
        productRoutes.POST("/", controllers.CreateProduct)
        productRoutes.GET("/", controllers.GetProducts)
        productRoutes.GET("/:id", controllers.GetProductByID)
        productRoutes.PUT("/:id", controllers.UpdateProduct)
        productRoutes.DELETE("/:id", controllers.DeleteProduct)
        productRoutes.POST("/:id/restore", controllers.RestoreProduct)
    }

    return r
}
