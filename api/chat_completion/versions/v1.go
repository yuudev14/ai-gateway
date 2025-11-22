package chatcompletionv1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatCompletionController struct {
}

func NewChatCompletionController() *ChatCompletionController {
	return &ChatCompletionController{}
}

func (c *ChatCompletionController) ChatCompletion(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		ctx.String(500, "Streaming not supported")
		return
	}

	for i := 1; i <= 5; i++ {
		ctx.Writer.Write([]byte("data: Hello " + time.Now().Format(time.RFC3339) + "\n\n"))
		flusher.Flush()
		time.Sleep(1 * time.Second)
	}

}
