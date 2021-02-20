package intent

import "context"

//DefaultHandler is the default implementation of the intent request Handler
type DefaultHandler struct {
}

// Build will extract the statement into its constitutent parts
// and return the request
func (d *DefaultHandler) Build(ctx context.Context, statement string) (Request, error) {
	return Request{}, nil
}

var handler = &DefaultHandler{}

// Handle is a wrapper around DefaultHandler to build a request
// from the provided statement
func Handle(ctx context.Context, statement string) (Request, error) {
	return handler.Build(ctx, statement)
}
