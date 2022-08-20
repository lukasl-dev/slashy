package slashy

type Command struct {
	// Runner runs the Command within a given Context. Every Command must have
	// at least a Runner.
	Runner Runner `json:"runner,omitempty"`

	// AutoCompleter provides auto-completion for the Command. This is optional.
	AutoCompleter AutoCompleter `json:"autoCompleter,omitempty"`

	// ErrorResponder is the command's individual ErrorResponder used to
	// construct interaction responses from errors. If this is nil, the Router's
	// ErrorResponder is used.
	ErrorResponder ErrorResponder `json:"errorResponder,omitempty"`
}
