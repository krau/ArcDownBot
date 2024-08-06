package bot

import (
	"arcdownbot/config"
	"fmt"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

var (
	Bot               *telego.Bot
	MainChannelChatID telego.ChatID
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
	MainChannelChatID = telegoutil.Username(config.Cfg.Usernames[0])
	return nil
}
