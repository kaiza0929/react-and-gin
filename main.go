package main

import (
	"gin_practice/controller"
	"gin_practice/service"

	"github.com/gin-gonic/gin"
)

func main() {

	var initService = service.InitService{}
	initService.Init()

	var engine = gin.Default()

	var apiEngine = engine.Group("/api")
	{
		/* localhost:8000/api/logにアクセスした場合のルーティング */
		var testLogApiEngine = apiEngine.Group("/log")
		{
			testLogApiEngine.GET("/", controller.FindTestLog)
		}

		/* localhost:8000/api/suggestionにアクセスした場合のルーティング */
		var testSuggestionApiEngine = engine.Group("/suggestion")
		{
			testSuggestionApiEngine.POST("/", controller.CreateTestSuggestion)
		}

	}

	engine.Run(":8000")

}
