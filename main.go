package main

import (
	"go-openai/openai"
)

func main1() {
	//os.Setenv("OPENAI_APIKEY", "inSGSo58WBtjZDt6D6SmT3BlbkFJzYTGWVxr6OfheP8hRQ18")
	//os.Setenv("GO_APIKEY", "aiuuxia")
	//os.Setenv("SERVER_PORT", "8082")
	openai.NewOpenAI().CreateOpenAiProxy()
	//TestProxy()
}
