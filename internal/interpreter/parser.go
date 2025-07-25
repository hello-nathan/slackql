package interpreter

import (
	"errors"
	"log"
	"strings"
)

type Command struct {
	Channel   string // collection/table
	Operation string // find/insert/etc.
	Query     string //
}

// @slackql command
func ParseText(text string) (Command, error) {
	parts := strings.Split(text, " ")
	if len(parts) != 2 {
		return Command{}, errors.New("invalid request")
	}
	return parseCommand(parts[1])
}

// channel.operation(query)
func parseCommand(command string) (Command, error) {
	parts := strings.Split(command, ".")
	if len(parts) != 2 {
		log.Printf("invalid command: %v", command)
		return Command{}, errors.New("invalid command")
	}
	channel := parts[0]
	idx := strings.IndexRune(parts[1], '(')
	if idx == -1 || !strings.HasSuffix(parts[1], ")") {
		panic("invalid format")
	}
	operation := parts[1][:idx]
	query := parts[1][idx+1 : len(parts[1])-1]

	// FIXME check valid operation and query

	return Command{
		Channel:   channel,
		Operation: operation,
		Query:     query,
	}, nil
}
