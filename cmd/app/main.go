package main

import (
	"context"
	"core/bot/cmd/handlers"
	"core/bot/cmd/middlewares"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatalf("Error while loading env file with %s", err)
	}
}

func main() {
	log.Println("Starting bot...")

	botToken := os.Getenv("TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
		bot.WithMiddlewares(middlewares.ShowMessageWithUserID, middlewares.ShowMessageWithUserName),
		// bot.WithDebug(),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}
	b.Start(ctx)
}
