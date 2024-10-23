package routes

import (
	"learn/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        productRoutes := api.Group("/products")
        {
            productRoutes.POST("/", controllers.CreateProduct)
            productRoutes.GET("/", controllers.GetProducts)
            productRoutes.GET("/:id", controllers.GetProductByID)
            productRoutes.PUT("/:id", controllers.UpdateProduct)
            productRoutes.DELETE("/:id", controllers.DeleteProduct)
            productRoutes.POST("/:id/restore", controllers.RestoreProduct)
        }
    }

    return r
}
