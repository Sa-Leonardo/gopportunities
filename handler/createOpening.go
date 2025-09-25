package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Sa-Leonardo/gopportunities/schemas"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	// Faz o bind do JSON para popular o meu request
	// ctx.BindJSON(&request)

	//Faz o bind do JSON para popular o meu request e já faz a verificação erro de formato do body
	if err := ctx.BindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "malformed request body"})
        return
    }

	if err := request.Validate(); err != nil {
		logger.Error("validate error:", err.Error()) // loga no terminal
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening %v", err.Error())
		sendError(ctx,http.StatusInternalServerError, "error creating opening on database")
		return
	}
	
	sendSuccess(ctx, "create-opening", opening)
}
