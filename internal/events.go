package internal

import (
	"log"
	"strings"

	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
)

func (bot *Bot) onMessage(session *dg.Session, msg *dg.MessageCreate) {
	// Ignore bot messages
	if msg.Author.Bot {
		return
	}

	msgAuthor := msg.Author.ID
	if len(msg.Attachments) == 0 {
		// mention the posted message in the NotificationChannel
		botMessage := "Posted a message without an attachment. \n Message Content: " + msg.Content
		util.Mention(session, msgAuthor, bot.Config.NotificationChannel, botMessage)
		// reply to the posted message
		util.Reply(bot.Client, msg.Message, "Posts without an attachment are not allowed in this channel")
		// remove the message
		session.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}

	// we can ask the authentication server if they're an admin of the bot
	// isAdmin defaults to "false" if err is not nil
	isAdmin, err := bot.Auth.CheckMember(msg.Author.ID)

	// let's split their message up into arguments
	// args = [prefix, sub-command name]
	args := strings.Fields(msg.Content)

	if len(args) < 2 { // this would mean args is [prefix] which at that point ignore them
		return
	}

	// we can now find out what command they were calling
	switch args[1] {
	case "help":
		// if all your commands are admin-relate then just wrap the whole switch statement
		// with this if statement

		// you can also just do "if isAdmin", the error doesn't matter that much if you want cleaner
		// code.
		if err != nil {
			util.Reply(bot.Client, msg.Message, "Failed to contact auth server")
		} else if isAdmin {
			bot.helpMessage(msg.Message)
		} else {
			util.Reply(bot.Client, msg.Message, "You must be an admin to run this command.")
		}
	}
}

func (bot *Bot) onReady(_ *dg.Session, ready *dg.Ready) {
	log.Printf("client bot - ready as %s#%s", ready.User.Username, ready.User.Discriminator)
}
