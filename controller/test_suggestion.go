package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTestSuggestion(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": "ボタン連打はいかがでしょうか",
	})
}
