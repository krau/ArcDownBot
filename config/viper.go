package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Token    string `toml:"token" mapstructure:"token"`
	BotApi   string `toml:"bot_api" mapstructure:"bot_api"`
	Interval int    `toml:"interval" mapstructure:"interval"`
	CacheDir string `toml:"cache_dir" mapstructure:"cache_dir"`
	Username string `toml:"username" mapstructure:"username"`
	ChatID   int64  `toml:"chat_id" mapstructure:"chat_id"`
}

var Cfg *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	viper.SetDefault("bot_api", "https://api.telegram.org")
	viper.SetDefault("cache_dir", "./cache")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file")
		os.Exit(1)
	}
	Cfg = &Config{}
	if err := viper.Unmarshal(Cfg); err != nil {
		fmt.Println("Error unmarshalling config")
		os.Exit(1)
	}
}
