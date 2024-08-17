package handlers

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	err, status := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Я пока ничего не умею",
	})

	if err == nil {
		log.Println(status)
	}
}
