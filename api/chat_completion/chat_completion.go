package chatcompletion_api

import (
	"github.com/gin-gonic/gin"
	chatcompletionv1 "github.com/yuudev14/ai-gateway/api/chat_completion/versions"
)

func SetupChatCompletionController(app *gin.RouterGroup) {
	controller := chatcompletionv1.NewChatCompletionController()
	app.POST("/v1/chat/completions", controller.ChatCompletion)
}
