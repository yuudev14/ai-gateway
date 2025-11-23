package llm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuudev14/ai-gateway/env"
)

func RequestToLlm(reqBody string) (*http.Response, error) {
	url := env.Settings.LLM_BASE_URL + "/v1/chat/completions"
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func ChatCompletion(ctx *gin.Context, reqBody string) error {
	resp, respErr := RequestToLlm(reqBody)

	if respErr != nil {
		return respErr
	}
	defer resp.Body.Close()
	respBodyByte, _ := io.ReadAll(resp.Body)
	var respBody interface{}
	if err := json.Unmarshal(respBodyByte, &respBody); err != nil {
		return err
	}
	fmt.Println(resp.Header)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.IndentedJSON(resp.StatusCode, respBody)
	return nil
}

func StreamChatCompletion(ctx *gin.Context, reqBody string) error {
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	flusher, ok := ctx.Writer.(http.Flusher)
	if !ok {
		ctx.String(500, "Streaming not supported")
		return nil
	}

	resp, respErr := RequestToLlm(reqBody)

	if respErr != nil {
		return respErr
	}
	defer resp.Body.Close()

	fmt.Println(resp.Header)
	scanner := bufio.NewScanner(resp.Body)
	var re = regexp.MustCompile(`^data: `)
	for scanner.Scan() {
		line := scanner.Text()
		formattedString := re.ReplaceAllString(line, "")
		fmt.Println(formattedString)
		ctx.Writer.Write([]byte(line + "\n"))
		flusher.Flush()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
