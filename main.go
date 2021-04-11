package main

import (
	"context"
	"log"
	"os"
	"errors"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {

	err2 := godotenv.Load()
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	slackbotToken := os.Getenv("SLACKBOT_TOKEN")

	bot := slacker.NewClient(slackbotToken, slacker.WithDebug(true))

	messageReplyDefinition := &slacker.CommandDefinition{
		Description: "Tests errors in new messages",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.ReportError(errors.New("Oops!"))
		},
	}

	threadReplyDefinition := &slacker.CommandDefinition{
		Description: "Tests errors in threads",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.ReportError(errors.New("Oops!"), slacker.WithThreadError(true))
		},
	}

	bot.Command("message", messageReplyDefinition)
	bot.Command("thread", threadReplyDefinition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
