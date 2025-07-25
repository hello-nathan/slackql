package main

import (
	"fmt"
	"log"

	"github.com/hello-nathan/slackql/internal/interpreter"
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
		switch msg.Type {
		case socketmode.EventTypeEventsAPI:
			ev, ok := msg.Data.(slackevents.EventsAPIEvent)
			if !ok {
				log.Printf("Failed to parse event: %v", msg.Data)
				continue
			}
			h.client.Ack(*msg.Request)
			switch innerEvent := ev.InnerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:

				command, err := interpreter.ParseText(innerEvent.Text)
				if err != nil {
					log.Printf("Failed to parse request: %v", err)
					continue
				}
				temp := fmt.Sprintf("%v : %v : %v", command.Channel, command.Operation, command.Query)
				_, _, err = h.api.PostMessage(
					innerEvent.Channel,
					slack.MsgOptionText(temp, false))
				if err != nil {
					log.Printf("Failed to post message: %v", err)
				}
			default:
				log.Printf("Unexpected inner event: %v", innerEvent)
			}
		default:
			log.Printf("Unexpected message: %v", msg)
		}
	}
}
