package main

import (
	"log"
	"time"

	"github.com/lrstanley/girc"
)

/*
Configure and run the bot

Args:

	bot: The bot to run
	config: Configuration for the bot
*/
func HandleBot(bot *girc.Client, config *BotConfig) error {
	log.Println("Setting up up", config.Name)
	var posts []string

	bot.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		if config.Onconnect != "" {
			err := c.Cmd.SendRaw(config.Onconnect)
			if err != nil {
				log.Println("Error sending onconnect command:", err)
				return
			}
		}

		c.Cmd.Join(config.Channel)
	})

	bot.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		for {
			log.Printf("Parsing RSS (%s)\n", config.RssURL)

			msg, err := ParseRss(config)
			if err != nil {
				log.Printf("Failed to format RSS (%s): %s\n", config.RssURL, err)
			}

			if len(msg) == 0 {
				log.Printf("No messages found for RSS (%s)\n", config.RssURL)
			}

			for _, m := range msg {
				if m == "" || inArray(posts, m){
					continue
				}
				c.Cmd.Message(config.Channel, girc.Fmt(m))
				for(len(posts)-1 >= config.MaxBacklog) {
					posts = posts[1:]
				}
				posts = append(posts, m)
				time.Sleep(time.Duration(config.RssAntiFlood) * time.Second)
			}

			time.Sleep(time.Duration(config.RssFetchInterval) * time.Second)
		}
	})

	return bot.Connect()
}
