package slashy

import "github.com/bwmarrin/discordgo"

type AutoCompleter interface {
	// AutoComplete returns an array of choices that match the current input of
	// the Context. The focused option is used to determine which option is
	// currently focused and, accordingly, the choices should be filtered.
	AutoComplete(ctx *Context, focused *discordgo.ApplicationCommandInteractionDataOption) []*discordgo.ApplicationCommandOptionChoice
}

// AutoCompleterFunc is a function that implements the AutoCompleter interface.
type AutoCompleterFunc func(ctx *Context, focused *discordgo.ApplicationCommandInteractionDataOption) []*discordgo.ApplicationCommandOptionChoice

// AutoComplete calls fn() itself and returns the result.
func (fn AutoCompleterFunc) AutoComplete(ctx *Context, focused *discordgo.ApplicationCommandInteractionDataOption) []*discordgo.ApplicationCommandOptionChoice {
	return fn(ctx, focused)
}
