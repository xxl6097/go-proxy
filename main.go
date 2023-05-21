package main

import (
	"go-openai/openai"
	"os"
)

func main() {
	os.Setenv("OPENAI_APIKEY", "inSGSo58WBtjZDt6D6SmT3BlbkFJzYTGWVxr6OfheP8hRQ18")
	os.Setenv("GO_APIKEY", "clife-001")
	os.Setenv("SERVER_PORT", "8080")
	openai.NewOpenAI().CreateOpenAiProxy()
}
