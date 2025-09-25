package router

import (
	docs "github.com/Sa-Leonardo/gopportunities/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"github.com/Sa-Leonardo/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
)

// TUDO QUE ESTÁ NO MESMO PACKAGE É ACESSIVEL EM OUTROS ARQUIVOS


func initializeRoutes(router *gin.Engine) {

	// Initialize handler
	handler.InitializeHandler()
	BasePath := "/api/v1"
	docs.SwaggerInfo.BasePath = BasePath

	var v1 *gin.RouterGroup = router.Group("/api/v1")
	

	{
		//defined routes
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	//Initialize Swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	
	// - http://localhost:8080/swagger/index.html
}