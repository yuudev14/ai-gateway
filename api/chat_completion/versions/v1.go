package chatcompletionv1

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/yuudev14/ai-gateway/internal/llm"
)

type ChatCompletionController struct {
}

func NewChatCompletionController() *ChatCompletionController {
	return &ChatCompletionController{}
}

func (c *ChatCompletionController) ChatCompletionProxy(ctx *gin.Context) {
	var data map[string]interface{}
	reqBody, _ := io.ReadAll(ctx.Request.Body)

	if err := json.Unmarshal(reqBody, &data); err != nil {
		panic(err)
	}
	stream, ok := data["stream"]
	if !ok {
		stream = false
	}

	if stream == true {
		llm.StreamChatCompletion(ctx, string(reqBody))
		return
	}
	llm.ChatCompletion(ctx, string(reqBody))
}
