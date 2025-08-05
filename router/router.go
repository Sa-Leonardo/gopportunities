package router

	// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo (gin	// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo, nesse caso (gin

	import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize router
	router := gin.Default()

  //Initialize Routes
  initializeRoutes(router)



  // Run the server
  router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
