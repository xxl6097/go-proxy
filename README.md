#go-openai
https://www.cnblogs.com/kenx/
http://openai.soboys.cn/login
https://apifox.com/apidoc/shared-d877c066-8cb9-44a3-bbc6-f38174766ae9
18.140.235.243
https://www.cnblogs.com/kenx/p/17408763.html?share_token=C7D47CBB-AFFB-4985-A77D-B4823B0EF8CE&tt_from=copy_link&utm_source=copy_link&utm_medium=toutiao_ios&utm_campaign=client_share%20ChatGPT4%E9%80%9A%E9%81%93%E5%BC%80%E6%94%BE%E6%8E%A5%E5%85%A5%E5%9F%BA%E4%BA%8EOPEN%20AI%20%E5%B9%B3%E5%8F%B0%E4%BD%A0%E7%9A%84%E4%BB%BB%E4%BD%95APP%20%E5%8F%AF%E4%B8%80%E9%94%AE%E6%8E%A5%E5%85%A5AI%20-%20%E4%BB%8A%E6%97%A5%E5%A4%B4%E6%9D%A1#chatgpt4-%E5%85%8D%E8%B4%B9%E6%94%AF%E6%8C%81
```shell
curl http://18.140.235.243:9090/v1/chat/completions \
    -H "Content-Type: application/json" \
    -H "Go-Authorization: aiuuxia" \
    -d '{
       "model": "gpt-4-0314",
       "messages": [{"role": "user", "content": "今天天气很好，请以天气做一首诗!"}],
       "temperature": 0.7
     }'

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