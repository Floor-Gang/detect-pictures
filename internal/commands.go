package internal

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
)

// Example command
func (bot *Bot) helpMessage(msg *dg.Message) {
	util.Reply(bot.Client, msg, "Looks for pictures in posts. If a post doesn't have a picture attachment, it reports it in another channel (NotificationChannel in config).\nThen deletes the posted messages wiht a reply stating text only messages are not allowed.")
}
