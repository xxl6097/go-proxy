package main

import "go-openai/gproxy"

func main() {
	//os.Setenv("OPENAI_APIKEY", "inSGSo58WBtjZDt6D6SmT3BlbkFJzYTGWVxr6OfheP8hRQ18")
	//os.Setenv("GO_APIKEY", "aiuuxia")
	//os.Setenv("SERVER_PORT", "8081")
	apikey, keys, port := gproxy.ObtainEnv()
	gproxy.CreateOpenAiDefaultProxy(port, apikey, keys)
}
