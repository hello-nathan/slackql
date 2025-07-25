package main

import (
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	botToken := getEnv("SLACK_BOT_TOKEN", "")
	appToken := getEnv("SLACK_APP_TOKEN", "")
	api := slack.New(botToken, slack.OptionAppLevelToken(appToken)) // FIXME options on api
	client := socketmode.New(api)
	handler := &Handler{
		client: client,
		api:    api,
	}
	go handler.Handle()
	client.Run()
}

func getEnv(key string, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
