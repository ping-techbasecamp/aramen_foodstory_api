package main

import (
	"aramen-api/cmd/api/di"
	"aramen-api/srv/logs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logs.SetupLogger()
	server, cleanupFn, err := di.InitializeApplication()
	defer cleanupFn()
	if err != nil {
		logs.Error().Err(err).Msg("Cannot initialize application")
		panic(err)
	}

	server.Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
