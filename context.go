package slashy

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Context struct {
	*discordgo.InteractionCreate

	// Command is the command name that was used to invoke the interaction.
	Command string

	// Session is the session that received the interaction-create event. It can
	// be used to interact directly with Discord.
	Session *discordgo.Session
}

// newContext returns a new Context for the given interaction create event.
func newContext(cmd string, ses *discordgo.Session, evt *discordgo.InteractionCreate) *Context {
	return &Context{
		Command:           cmd,
		Session:           ses,
		InteractionCreate: evt,
	}
}

// Option searches for the first option that matches the given name and returns
// it. If no option is found, nil is returned instead
//
// Option names are case-insensitive. Therefore, "option" and "OPTION" are
// considered the same.
func (ctx *Context) Option(name string) *discordgo.ApplicationCommandInteractionDataOption {
	opts := ctx.ApplicationCommandData().Options
	for _, opt := range opts {
		if strings.EqualFold(opt.Name, name) {
			return opt
		}
	}
	return nil
}

// FocusedOption searches for the first option that is focused and returns it. If
// no option is found, nil is returned instead.
func (ctx *Context) FocusedOption() *discordgo.ApplicationCommandInteractionDataOption {
	opts := ctx.ApplicationCommandData().Options
	for _, opt := range opts {
		if opt.Focused {
			return opt
		}
	}
	return nil
}
