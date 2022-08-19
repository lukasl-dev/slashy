package slashy

import "github.com/bwmarrin/discordgo"

type AutoCompleter interface {
	// AutoComplete returns an array of choices that match the current input of
	// the Context.
	AutoComplete(ctx *Context) ([]*discordgo.ApplicationCommandOptionChoice, error)
}
