package slashy

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Context struct {
	*discordgo.InteractionCreate
}

// newContext returns a new Context for the given interaction create event.
func newContext(evt *discordgo.InteractionCreate) *Context {
	return &Context{InteractionCreate: evt}
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
