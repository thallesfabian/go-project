package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// initialize routes
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETAR Opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Openings",
			})
		})
	}
}

