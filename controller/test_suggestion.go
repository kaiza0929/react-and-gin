package controller

import (
	"gin_practice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testSuggetionService = service.TestSuggestionService{}

func CreateTestSuggestion(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": testSuggetionService.CreateTestSuggestion(),
	})
}

func FindTestSuggestions(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"suggestions": testSuggetionService.FindTestSuggestions(),
	})
}
