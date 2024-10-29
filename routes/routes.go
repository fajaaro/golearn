package routes

import (
	"learn/app/controllers"
	"learn/app/middleware"

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
        productRoutes.POST("/", middleware.CheckPermission("create product"), controllers.CreateProduct)
        productRoutes.GET("/", middleware.CheckPermission("view product"), controllers.GetProducts)
        productRoutes.GET("/:id", middleware.CheckPermission("view product"), controllers.GetProductByID)
        productRoutes.PUT("/:id", middleware.CheckPermission("update product"), controllers.UpdateProduct)
        productRoutes.DELETE("/:id", middleware.CheckPermission("delete product"), controllers.DeleteProduct)
        productRoutes.POST("/:id/restore", middleware.CheckPermission("restore product"), controllers.RestoreProduct)

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
        permissionRoutes.Use(middleware.CheckSuperAdmin())
        permissionRoutes.POST("/", controllers.CreatePermission)
        permissionRoutes.GET("/", controllers.GetPermissions)
        permissionRoutes.GET("/:id", controllers.GetPermissionByID)
        permissionRoutes.PUT("/:id", controllers.UpdatePermission)
        permissionRoutes.DELETE("/:id", controllers.DeletePermission)
        permissionRoutes.POST("/:id/restore", controllers.RestorePermission)

        rolePermissionRoutes := api.Group("/role-permissions")
        rolePermissionRoutes.Use(middleware.JWTMiddleware())
        rolePermissionRoutes.Use(middleware.CheckSuperAdmin())
        rolePermissionRoutes.POST("/", controllers.CreateRolePermission)
        rolePermissionRoutes.GET("/", controllers.GetRolePermissions)
        rolePermissionRoutes.GET("/:id", controllers.GetRolePermissionByID)
        rolePermissionRoutes.PUT("/:id", controllers.UpdateRolePermission)
        rolePermissionRoutes.DELETE("/:id", controllers.DeleteRolePermission)

        userPermissionRoutes := api.Group("/user-permissions")
        userPermissionRoutes.Use(middleware.JWTMiddleware())
        userPermissionRoutes.Use(middleware.CheckSuperAdmin())
        userPermissionRoutes.POST("/", controllers.CreateUserPermission)
        userPermissionRoutes.GET("/", controllers.GetUserPermissions)
        userPermissionRoutes.GET("/:id", controllers.GetUserPermissionByID)
        userPermissionRoutes.PUT("/:id", controllers.UpdateUserPermission)
        userPermissionRoutes.DELETE("/:id", controllers.DeleteUserPermission)

        userRoleRoutes := api.Group("/user-roles")
        userRoleRoutes.Use(middleware.JWTMiddleware())
        userRoleRoutes.Use(middleware.CheckSuperAdmin())
        userRoleRoutes.POST("/", controllers.CreateUserRole)
        userRoleRoutes.GET("/", controllers.GetUserRoles)
        userRoleRoutes.GET("/:id", controllers.GetUserRoleByID)
        userRoleRoutes.PUT("/:id", controllers.UpdateUserRole)
        userRoleRoutes.DELETE("/:id", controllers.DeleteUserRole)
    }

    return r
}
