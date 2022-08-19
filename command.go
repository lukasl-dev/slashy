package slashy

type Command struct {
	// Runner runs the Command within a given Context. Every Command must have
	// at least a Runner.
	Runner Runner `json:"runner,omitempty"`

	// AutoCompleter provides auto-completion for the Command. This is optional.
	AutoCompleter AutoCompleter `json:"autoCompleter,omitempty"`
}
