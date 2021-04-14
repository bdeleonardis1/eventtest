package events

func StartListening(port string, envVarName string) {
	if port == "" {
		port = "1111"
	}

	if envVarName == "" {
		envVarName = "EVENTTESTPORT"
	}

	// Since this is in a goroutine, we don't need to worry about
	// stopping the server from listening. When the program terminates
	// this will stop.
	go createServer()

	// TODO: setup environment variable
}

func StopListening() {
	// TODO: cleanup environment variable.
}
