package bot

import (
	"arcdownbot/config"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

var (
	Bot           *telego.Bot
	ChannelChatID telego.ChatID
)

func InitBot() error {
	fmt.Println("Bot initing...")
	var err error
	Bot, err = telego.NewBot(
		config.Cfg.Token,
		telego.WithDefaultLogger(false, true),
		telego.WithAPIServer(config.Cfg.BotApi),
	)
	if err != nil {
		return err
	}
	if config.Cfg.Username != "" {
		ChannelChatID = telegoutil.Username(config.Cfg.Username)
	} else {
		ChannelChatID = telegoutil.ID(config.Cfg.ChatID)
	}
	return nil
}

func Run() {
	if Bot == nil {
		fmt.Println("Bot not inited")
		os.Exit(1)
	}
	fmt.Println("Start Bot")
	updates, err := Bot.UpdatesViaLongPolling(
		&telego.GetUpdatesParams{
			Offset: -1,
			AllowedUpdates: []string{
				telego.MessageUpdates,
				telego.ChannelPostUpdates,
				telego.CallbackQueryUpdates,
				telego.InlineQueryUpdates,
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	botHandler, err := telegohandler.NewBotHandler(Bot, updates)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer botHandler.Stop()
	defer Bot.StopLongPolling()

	botHandler.Use(telegohandler.PanicRecovery())

	botHandler.Start()
}
