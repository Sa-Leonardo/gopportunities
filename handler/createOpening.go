package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := struct {
		Role string `json:"role"`
	}{}

	// Faz o bind do JSON
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("erro no bind JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loga o request recebido
	logger.Infof("request received: %+v", request)

	// Retorna uma resposta JSON
	ctx.JSON(http.StatusOK, gin.H{
		"message": "vaga criada com sucesso",
		"data":    request,
	})
}
