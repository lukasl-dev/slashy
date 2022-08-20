package slashy

type CommandRoute struct {
	// Runner runs the CommandRoute within a given Context. Every CommandRoute must have
	// at least a Runner.
	Runner Runner `json:"runner,omitempty"`

	// AutoCompleter provides auto-completion for the CommandRoute. This is optional.
	AutoCompleter AutoCompleter `json:"autoCompleter,omitempty"`

	// ErrorResponder is the command's individual ErrorResponder used to
	// construct interaction responses from errors. If this is nil, the Router's
	// ErrorResponder is used.
	ErrorResponder ErrorResponder `json:"errorResponder,omitempty"`
}
