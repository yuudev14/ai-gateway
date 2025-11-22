package main

import (
	"github.com/yuudev14/ai-gateway/api"
	"github.com/yuudev14/ai-gateway/env"
)

func main() {
	env.Setup()
	app := api.InitRouter()

	app.Run(":9999") // default: :8080
}
