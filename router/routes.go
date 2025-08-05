package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TUDO QUE ESTÁ NO MESMO PACKAGE É ACESSIVEL EM OUTROS ARQUIVOS


func initializeRoutes(router *gin.Engine) {

	var v1 *gin.RouterGroup = router.Group("/api/v1")
	
	var err error
	
	{
		//defined routes
		v1.GET("/opening", func(ctx *gin.Context)  {
			ctx.JSON(http.StatusOK, gin.H{
				"msg" : "GET opening",
			})
			
		})

		v1.POST("/opening", func(ctx *gin.Context)  {
			ctx.JSON(http.StatusOK, gin.H{
				"msg" : "POST opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context)  {
			ctx.JSON(http.StatusOK, gin.H{
				"msg" : "DELETE opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context)  {
			ctx.JSON(http.StatusOK, gin.H{
				"msg" : "PUT opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context)  {

			if err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"msg" : "GET openings",
			})

			
		})
	}

}