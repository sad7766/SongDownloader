package main

import (
    "log"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


func main() {
    // Replace YOUR_BOT_TOKEN with the token you obtained from BotFather
    bot, err := tgbotapi.NewBotAPI("6624535991:AAFDyQx_eIqam62I5ZYm6dn0ahaw_U1fT3g")
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = true // Uncomment this line to see debug logs

    log.Printf("Authorized on account %s", bot.Self.UserName)

    // Set up handlers for various update types
    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 30

    updates, err := bot.GetUpdatesChan(updateConfig)
    if err != nil {
        log.Panic(err)
    }

    for update := range updates {
        if update.Message == nil {
            continue
        }

        // Handle incoming messages, commands, etc.
        // Example: Reply to messages that start with "/hello"
        if update.Message.Text == "/hello" {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, world!")
            bot.Send(msg)
        }
    }
}