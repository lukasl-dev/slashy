package slashy

type ErrorResponder interface {
	// RespondError mutates the given Response according to the given error.
	RespondError(ctx *Context, resp *Response, err error)
}

// ErrorResponderFunc is a function that implements the ErrorResponder interface.
type ErrorResponderFunc func(ctx *Context, resp *Response, err error)

// RespondError calls fn() itself and returns the result.
func (fn ErrorResponderFunc) RespondError(ctx *Context, resp *Response, err error) {
	fn(ctx, resp, err)
}

// defaultErrorResponder is the default ErrorResponder.
var defaultErrorResponder ErrorResponder = ErrorResponderFunc(func(ctx *Context, resp *Response, err error) {
	resp.Content("An error occurred: " + err.Error())
})
