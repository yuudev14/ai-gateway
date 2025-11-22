package chatcompletionv1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuudev14/ai-gateway/env"
)

type ChatCompletionController struct {
}

func NewChatCompletionController() *ChatCompletionController {
	return &ChatCompletionController{}
}

func (c *ChatCompletionController) ChatCompletion(ctx *gin.Context) {
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
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")

		flusher, ok := ctx.Writer.(http.Flusher)
		if !ok {
			ctx.String(500, "Streaming not supported")
			return
		}

		url := env.Settings.LLM_BASE_URL + "/v1/chat/completions"

		req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read streaming response line by line
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			ctx.Writer.Write([]byte(line + "\n"))
			fmt.Println(line)
			flusher.Flush()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
		return
	}

	url := env.Settings.LLM_BASE_URL + "/v1/chat/completions"

	req, err := http.NewRequest("POST", url, strings.NewReader(string(reqBody)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBodyByte, _ := io.ReadAll(resp.Body)
	var respBody interface{}
	if err := json.Unmarshal(respBodyByte, &respBody); err != nil {
		panic(err)
	}
	fmt.Println(resp.Header)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.IndentedJSON(resp.StatusCode, respBody)

}
