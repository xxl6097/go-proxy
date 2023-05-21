package main

import (
	"go-openai/openai"
)

func main() {
	openai.NewOpenAI().CreateOpenAiProxy()
}
