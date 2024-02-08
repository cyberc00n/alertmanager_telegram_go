### Telegram alertmanager webhook

#### 📦 Requirements:
- go 1.21.1
- web-server for routing
- or you can deploy it as a container


#### 🐋 Docker deploy:

1. Clone this repo:
```
git clone https://github.com/cyberc00n/alertmanager_telegram_webhook.git
```
2. Build image:

```
docker build -t alertmanager-telegram-bot:latest .
```

3. Run service (replace token with your telegram bot token and chat id with your channel id)

```
docker run -d -e TELEGRAM_BOT_TOKEN=token \
  -e TELEGRAM_CHAT_ID=chat_id \
  -p 8080:8080 \
  --name telegram_alertmanager \
 alertmanager-telegram-bot:latest

 ```

#### 👨🏻‍💻 Usage

Place this snippet in the receiver section of the alertmanager config.yml file:
 
 ```
receivers:
- name: 'telegram-receiver'
  webhook_configs:
  - url: 'http://ip_address:8080/alert'
    send_resolved: true
 ```

And change receiver names in routes
