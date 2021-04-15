package events

import (
	"context"
	"os"
)

const (
	envVarPortName = "EVENTTESTPORT"
)

type CloseableServer interface {
	Shutdown(context.Context) error
}

func StartListening(port string) CloseableServer {
	if port == "" {
		port = "1111"
	}
	os.Setenv(envVarPortName, port)

	server := createServer(port)

	os.Setenv(envVarPortName, port)
	return server
}

func StopListening(server CloseableServer) {
	err := server.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
	//os.Unsetenv(envVarPortName)
}
