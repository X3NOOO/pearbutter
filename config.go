package main

import (
	"os"

	"github.com/pelletier/go-toml"
)

type BotConfig struct {
	Name             string `toml:"name"`
	Address          string `toml:"address"`
	Password         string `toml:"password"`
	Port             int    `toml:"port"`
	Ssl              bool   `toml:"ssl"`
	Nick             string `toml:"nick"`
	User             string `toml:"user"`
	Realname         string `toml:"realname"`
	Channel          string `toml:"channel"`
	RssFetchInterval int    `toml:"rss_fetch_interval"`
	RssAntiFlood	 int    `toml:"rss_anti_flood"`
	RssURL           string `toml:"rss_url"`
	Formatting       string `toml:"formatting"`
	Onconnect        string `toml:"onconnect"`
}

type Config struct {
	Config struct {
		Logfile string `toml:"logfile"`
	} `toml:"config"`
	Servers []BotConfig `toml:"servers"`
}

func parseConfig(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = toml.Unmarshal(f, config)

	return config, err
}

func GetConfig(paths []string) (*Config, error) {
	var config *Config

	for _, path := range paths {
		fi, err := os.Stat(path)
		if err != nil {
			continue
		}

		if fi.IsDir() {
			continue
		}

		config, err = parseConfig(path)
		if err != nil {
			return nil, err
		}

		break
	}

	if config == nil {
		return nil, os.ErrNotExist
	}

	return config, nil
}
