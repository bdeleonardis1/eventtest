package events

import (
	"context"
)

type CloseableServer interface {
	Shutdown(context.Context) error
}

// StartListening must be called to start keeping track of events.
// It creates an http server listing at localhost:1111 and returns
// the server so that it can be shutdown.
func StartListening() CloseableServer {
	return createServer()
}

// StopListening shuts down the server. Callers of StartListening
// should immediately defer StopListening.
func StopListening(server CloseableServer) {
	err := server.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
