package slashy

import "github.com/bwmarrin/discordgo"

type Response struct {
	// response is the current response message. It is mutated by the methods
	// of this struct.
	response discordgo.InteractionResponseData
}

// newResponse returns a new Response.
func newResponse() *Response {
	return new(Response)
}

// TTS updates whether the response message should be sent as text-to-speech.
func (r *Response) TTS(tts bool) *Response {
	r.response.TTS = tts
	return r
}

// Content updates the content of the response message.
func (r *Response) Content(content string) *Response {
	r.response.Content = content
	return r
}

// Components updates the components of the response message.
func (r *Response) Components(components ...discordgo.MessageComponent) *Response {
	r.response.Components = components
	return r
}

// AddComponents adds the given components to the response message.
func (r *Response) AddComponents(components ...discordgo.MessageComponent) *Response {
	r.response.Components = append(r.response.Components, components...)
	return r
}

// Embeds updates the embeds of the response message.
func (r *Response) Embeds(embeds ...*discordgo.MessageEmbed) *Response {
	r.response.Embeds = embeds
	return r
}

// AddEmbeds adds the given embeds to the response message.
func (r *Response) AddEmbeds(embeds ...*discordgo.MessageEmbed) *Response {
	r.response.Embeds = append(r.response.Embeds, embeds...)
	return r
}

// AllowedMentions updates the allowed mentions of the response message.
func (r *Response) AllowedMentions(allowedMentions *discordgo.MessageAllowedMentions) *Response {
	r.response.AllowedMentions = allowedMentions
	return r
}

// Files updates the files of the response message.
func (r *Response) Files(files ...*discordgo.File) *Response {
	r.response.Files = files
	return r
}

// AddFiles adds the given files to the response message.
func (r *Response) AddFiles(files ...*discordgo.File) *Response {
	r.response.Files = append(r.response.Files, files...)
	return r
}

// Flags enables or disables the given flags on the response message.
func (r *Response) Flags(suppress, ephemeral bool) *Response {
	r.response.Flags = 0
	if suppress {
		r.response.Flags |= discordgo.MessageFlagsSuppressEmbeds
	}
	if ephemeral {
		r.response.Flags |= discordgo.MessageFlagsEphemeral
	}
	return r
}
