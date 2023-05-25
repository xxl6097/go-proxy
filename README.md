#go-openai

18.140.235.243

```shell

curl http://aispeaker.cc:9090/v1/models -H "Go-Authorization: aiuuxia"

curl http://aispeaker.cc:9090/v1/models \
  -H "Authorization: Bearer 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"

curl https://api.openai.com/v1/models \
  -H "Authorization: Bearer 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"


curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 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" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
  
  curl https://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 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" \
  -d '{
    "model": "gpt-4",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'


curl http://18.140.235.243/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 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" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
  
  curl http://api.openai.com/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer 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" \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "Hello!"}]
     }'

curl https://api.openai.com/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: clife-001" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'

curl http://18.140.235.243:9090/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: aiuuxia" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'
     
     
curl http://localhost:8080/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: clife-001" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'



curl http://openai.clife.net/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: aiclife" \
    -d '{
       "model": "gpt-3.5-turbo",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'

```