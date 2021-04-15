package events

import (
	"context"
)

type CloseableServer interface {
	Shutdown(context.Context) error
}

func StartListening() CloseableServer {
	return createServer()
}

func StopListening(server CloseableServer) {
	err := server.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
