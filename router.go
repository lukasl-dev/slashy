package slashy

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Router struct {
	// commands is a map of lower-case Command names to commands.
	commands map[string]*Command

	// errorResponder is the ErrorResponder that is used to construct error
	// responses.
	errorResponder ErrorResponder
}

// NewRouter returns a new router without any commands.
//
// The given ErrorResponder is used to construct error messages. If nil is
// given, the defaultErrorResponder is used.
func NewRouter(errResp ErrorResponder) *Router {
	if errResp == nil {
		errResp = defaultErrorResponder
	}

	return &Router{
		commands:       make(map[string]*Command),
		errorResponder: errResp,
	}
}

// Bind binds the given Command to the given name. If the name is already taken,
// the existing Command is overwritten.
//
// Command names are case-insensitive. Therefore, "command" and "COMMAND" are
// considered the same. Note that the name is stored in the lower-case.
func (r *Router) Bind(name string, cmd *Command) {
	switch {
	case name == "":
		panic("Bind(): name must not be empty")
	case cmd == nil:
		panic("Bind(): cmd must not be nil")
	case cmd.Runner == nil:
		panic("Bind(): cmd.Runner must not be nil")
	}

	r.put(name, cmd)
}

// AutoBind binds the given Command to the given name. If the name is already
// taken, the existing Command is overwritten.
//
// The difference between Bind and AutoBind is that AutoBind takes a Runner and
//  tests whether the Runner is also an AutoCompleter. If that is the case,
// the AutoCompleter is bound to the Command.
//
// Command names are case-insensitive. Therefore, "command" and "COMMAND" are
// considered the same. Note that the name is stored in the lower-case.
func (r *Router) AutoBind(name string, cmd Runner) {
	switch {
	case name == "":
		panic("AutoBind(): name must not be empty")
	case cmd == nil:
		panic("AutoBind(): cmd must not be nil")
	}

	com := r.get(name)
	if com == nil {
		com = new(Command)
	}
	com.Runner = cmd

	completer, isCompleter := cmd.(AutoCompleter)
	if isCompleter {
		com.AutoCompleter = completer
	}

	r.Bind(name, com)
}

// Route handles an interaction create events and routes it to the appropriate
// Command. Unknown interaction types and commands are ignored.
func (r *Router) Route(s *discordgo.Session, evt *discordgo.InteractionCreate) {
	ctx := newContext(evt)

	cmd := r.get(ctx.ApplicationCommandData().Name)
	if cmd == nil {
		return
	}

	resp := r.run(ctx, cmd)
	if resp == nil {
		return
	}

	_ = s.InteractionRespond(evt.Interaction, resp)
}

// handle handles an interaction create event and returns the response. If the
// interaction type is not supported, nil is returned.
func (r *Router) handle(ctx *Context, cmd *Command) *discordgo.InteractionResponse {
	switch ctx.Type {
	case discordgo.InteractionApplicationCommand:
		return r.run(ctx, cmd)
	case discordgo.InteractionApplicationCommandAutocomplete:
		return r.autoComplete(ctx, cmd)
	default:
		return nil
	}
}

// run handles an interaction command event and returns the response that
// includes the message data.
//
// Errors returned by the AutoCompleter are formatted property and are returned
// as an interaction response.
func (r *Router) run(ctx *Context, cmd *Command) *discordgo.InteractionResponse {
	resp := newResponse()

	err := cmd.Runner.Run(ctx, resp)
	if err != nil {
		r.errorResponder.RespondError(ctx, resp, err)
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &resp.response,
		}
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &resp.response,
	}
}

// autoComplete handles an interaction autocomplete event and returns the
// response that includes the choices. If the given cmd does not own an
// AutoCompleter, a response with a zero-length choices-slice is returned.
//
// Errors returned by the AutoCompleter are formatted property and are returned
// as an interaction response.
func (r *Router) autoComplete(ctx *Context, cmd *Command) *discordgo.InteractionResponse {
	if cmd.AutoCompleter == nil {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{
				Choices: make([]*discordgo.ApplicationCommandOptionChoice, 0),
			},
		}
	}

	choices := cmd.AutoCompleter.AutoComplete(ctx)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	}
}

// get returns the Command with the given name. If no Command with the given name
// exists, nil is returned.
//
// Command names are case-insensitive. Therefore, "command" and "COMMAND" are
// considered the same. Note that the name is stored in the lower-case.
func (r *Router) get(name string) *Command {
	return r.commands[strings.ToLower(name)]
}

// put inserts the given Command under the given name. If the name is already
// taken, the existing command is overwritten.
//
// Command names are case-insensitive. Therefore, "command" and "COMMAND" are
// considered the same. Note that the name is stored in the lower-case.
func (r *Router) put(name string, cmd *Command) {
	r.commands[strings.ToLower(name)] = cmd
}
