package slashy

import "github.com/bwmarrin/discordgo"

type AutoCompleter interface {
	// AutoComplete returns an array of choices that match the current input of
	// the Context.
	AutoComplete(ctx *Context) []*discordgo.ApplicationCommandOptionChoice
}

// AutoCompleterFunc is a function that implements the AutoCompleter interface.
type AutoCompleterFunc func(ctx *Context) []*discordgo.ApplicationCommandOptionChoice

// AutoComplete calls fn() itself and returns the result.
func (fn AutoCompleterFunc) AutoComplete(ctx *Context) []*discordgo.ApplicationCommandOptionChoice {
	return fn(ctx)
}
