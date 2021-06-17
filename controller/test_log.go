package controller

import (
	"gin_practice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindTestLog(context *gin.Context) {
	var testLogService = service.TestLogService{}
	context.JSON(http.StatusOK, gin.H{
		"logs": testLogService.FindTestLogs(),
	})
}
