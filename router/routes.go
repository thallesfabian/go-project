package router

import (

	"github.com/gin-gonic/gin"
	"github.com/thallesfabian/gooportunis.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	// initialize routes
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening",handler.ShowOpeningHandler)
		v1.POST("/opening" ,handler.CreateOpeningHandler)
		v1.DELETE("/opening",handler.DeleteOpeningHandler)
		v1.PUT("/opening",handler.UpdateOpeningHandler)
		v1.GET("/openings",handler.ListOpeningsHandler)
	}
}
