package mandatory

import "context"

type mandatoryContext struct{}

var contextKey = mandatoryContext{}

// Context function to create context with mandatory.
func Context(ctx context.Context, mandatoryRequest Request) context.Context {
	ctx = context.WithValue(ctx, contextKey, &mandatoryRequest)
	return ctx
}

// FromContext function to get mandatory from context.
func FromContext(ctx context.Context) Request {
	if nil == ctx {
		return Request{}
	}

	mandatory, found := ctx.Value(contextKey).(*Request)
	if !found || nil == mandatory {
		return Request{}
	}
	return *mandatory
}
