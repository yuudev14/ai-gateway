package main

import (
	"github.com/yuudev14/ai-gateway/api"
)

func main() {
	app := api.InitRouter()

	app.Run(":9999") // default: :8080
}
