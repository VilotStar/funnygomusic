package commands

import (
	"funnygomusic/bot"
	"funnygomusic/databaser"
	"funnygomusic/utils"
	"github.com/diamondburned/arikawa/v3/gateway"
	"strings"
)

func LastFmCommand(c *gateway.MessageCreateEvent, b *bot.Botter, args []string) {
	if len(args) == 0 {
		ListQueue(c, b, 0)
		return
	}
	switch strings.ToLower(args[0]) {
	case "register":
		{
			token := utils.GetIndex(args, 1)
			go b.Db.Save(databaser.AllowedUser{ID: uint64(c.Author.ID), Token: token})
		}
	}
}
