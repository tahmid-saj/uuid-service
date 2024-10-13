package routes

import (
	"net/http"
	"uuid-service/models"

	"github.com/gin-gonic/gin"
)

func getUUID(context *gin.Context) {
	uuid, err := models.GenerateUUID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate an UUID"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"response": uuid})
}