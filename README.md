#go-openai

18.140.235.243

```shell

curl http://18.140.235.243/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: clife-001" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'


```


```shell

curl http://localhost:8080/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: clife-001" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'


```