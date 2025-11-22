package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
	chatcompletion_api "github.com/yuudev14/ai-gateway/api/chat_completion"
)

func StartApi(app *gin.RouterGroup) {
	chatcompletion_api.SetupChatCompletionController(app)
}

func InitRouter() *gin.Engine {

	app := gin.Default()

	config := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}

	// Use CORS middleware
	app.Use(cors.New(config))

	docs.SwaggerInfo.BasePath = "./"

	apiGroup := app.Group("/api")

	StartApi(apiGroup)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app

}
