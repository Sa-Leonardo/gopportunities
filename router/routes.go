package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Sa-Leonardo/gopportunities/handler"
)

// TUDO QUE ESTÁ NO MESMO PACKAGE É ACESSIVEL EM OUTROS ARQUIVOS


func initializeRoutes(router *gin.Engine) {

	// Initialize handler
	handler.InitializeHandler()

	var v1 *gin.RouterGroup = router.Group("/api/v1")
	

	{
		//defined routes
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}