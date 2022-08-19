package slashy

type Runner interface {
	// Run executes the command with the given context. If the command fails,
	// the causing error is returned.
	//
	// The given Response is used to mutate the response message. The response
	// message is built after the command is run successfully.
	Run(ctx *Context, resp *Response) error
}
