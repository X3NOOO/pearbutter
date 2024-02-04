package main

import (
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
)

/*
Get all the new items since the last message from the RSS feed

Args:

	config: The bot configuration

Returns:

	[]string: The items in the RSS feed
*/
func ParseRss(config *BotConfig) ([]string, error) {
	doc, err := xmlquery.LoadURL(config.RssURL)
	if err != nil {
		return nil, err
	}

	formats := regexp.MustCompile(`%([\w>]*)%`).FindAllStringSubmatch(config.Formatting, -1)

	items := xmlquery.Find(doc, "//item")
	items = reverse(items)

	var posts []string
	for _, item := range items {
		var msg string = config.Formatting
		for _, f := range formats {
			msg = strings.ReplaceAll(msg, f[0], xmlquery.FindOne(item, f[1]).InnerText())
		}
		posts = append(posts, msg)
	}

	return posts, nil
}
