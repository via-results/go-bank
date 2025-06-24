package router

import "github.com/gin-gonic/gin"

// initRouter initializes the API routes for the application.
func initRouter(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/customers", func(ctx *gin.Context) {
			ctx.JSON(201, gin.H{"message": "Customer created successfully!"})
		})

		v1.GET("/customers/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Customer retrieved successfully!"})
		})

		v1.DELETE("/customers/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Customer deleted successfully!"})
		})

	}

}
