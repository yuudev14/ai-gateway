package chatcompletionv1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuudev14/ai-gateway/internal/llm"
	rest "github.com/yuudev14/ai-gateway/internal/rests"
)

type ChatCompletionController struct {
}

func NewChatCompletionController() *ChatCompletionController {
	return &ChatCompletionController{}
}

func (c *ChatCompletionController) ChatCompletionProxy(ctx *gin.Context) {
	var data map[string]interface{}
	response := rest.Response{C: ctx}
	reqBody, _ := io.ReadAll(ctx.Request.Body)

	if err := json.Unmarshal(reqBody, &data); err != nil {
		return
	}

	stream, ok := data["stream"]

	// if stream does not exist in the data, set stream to false
	if !ok {
		stream = false
	}

	if stream == true {
		err := llm.StreamChatCompletion(ctx, string(reqBody))
		if err != nil {
			response.ResponseError(http.StatusBadRequest, err.Error())
		}
		return
	}
	llm.ChatCompletion(ctx, string(reqBody))
}
