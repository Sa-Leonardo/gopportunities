package router

	// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo (gin	// sempre que u não der um nome para o import, ele vai ser importado pelo ultimo pathName do modulo, nesse caso (gin

	import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o router usando as configurações defalt do gin 
	router := gin.Default()

  // Definindo uma rota 
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  //Estamos rodando a API
  router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
