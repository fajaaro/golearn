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

        roleRoutes := api.Group("/roles")
        roleRoutes.Use(middleware.JWTMiddleware())
        roleRoutes.POST("/", controllers.CreateRole)
        roleRoutes.GET("/", controllers.GetRoles)
        roleRoutes.GET("/:id", controllers.GetRoleByID)
        roleRoutes.PUT("/:id", controllers.UpdateRole)
        roleRoutes.DELETE("/:id", controllers.DeleteRole)
        roleRoutes.POST("/:id/restore", controllers.RestoreRole)

        permissionRoutes := api.Group("/permissions")
        permissionRoutes.Use(middleware.JWTMiddleware())
        permissionRoutes.POST("/", controllers.CreatePermission)
        permissionRoutes.GET("/", controllers.GetPermissions)
        permissionRoutes.GET("/:id", controllers.GetPermissionByID)
        permissionRoutes.PUT("/:id", controllers.UpdatePermission)
        permissionRoutes.DELETE("/:id", controllers.DeletePermission)
        permissionRoutes.POST("/:id/restore", controllers.RestorePermission)
    }

    return r
}
