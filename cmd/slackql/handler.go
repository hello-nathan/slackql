package main

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type Handler struct {
	client *socketmode.Client
	api    *slack.Client
}

func (h *Handler) Handle() {
	for msg := range h.client.Events {
		switch ev := msg.Data.(type) {
		case *slackevents.MessageEvent:
			_, _, err := h.api.PostMessage(
				ev.Channel,
				slack.MsgOptionText("Hello, world!", false))
			if err != nil {
				log.Printf("Failed to post message: %v", err)
			}
		default:
			log.Printf("Unexpected message: %v", ev)
		}
	}
}
