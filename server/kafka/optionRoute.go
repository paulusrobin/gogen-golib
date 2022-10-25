package kafka

type registerRoutes struct {
	handlers map[string]HandlerFunc
}

// Apply function to implement Option.
func (r registerRoutes) Apply(o *options) {
	for topic, handler := range r.handlers {
		o.topics = append(o.topics, topic)
		o.routes[topic] = handler
	}
}

// RegisterRoute function to add route to server.
func RegisterRoute(topic string, handler HandlerFunc) Option {
	return registerRoutes{
		handlers: map[string]HandlerFunc{
			topic: handler,
		},
	}
}

// RegisterRoutes function to add multi route to server.
func RegisterRoutes(handlers map[string]HandlerFunc) Option {
	return registerRoutes{handlers: handlers}
}
