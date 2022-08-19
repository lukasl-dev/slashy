package slashy

type Runner interface {
	// Run executes the Command with the given context. If the Command fails,
	// the causing error is returned.
	//
	// The given Response is used to mutate the response message. The response
	// message is built after the Command is run successfully.
	Run(ctx *Context, resp *Response) error
}

// RunnerFunc is a function that implements the Runner interface.
type RunnerFunc func(ctx *Context, resp *Response) error

// Run calls fn() itself and returns the result.
func (fn RunnerFunc) Run(ctx *Context, resp *Response) error {
	return fn(ctx, resp)
}
