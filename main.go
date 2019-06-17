package main

import (
	"github.com/docker/go-plugins-helpers/sdk"
)

func main() {

	sdkHandler := sdk.NewHandler(`{"Implements": ["LoggingDriver"]}`)
	//Create handlers for startup and shutdown of the log driver
	sdkHandler.HandleFunc("/LogDriver.StartLogging", startLoggingHandler())
	sdkHandler.HandleFunc("/LogDriver.StopLogging", stopLoggingHandler())

	err := sdkHandler.ServeUnix("hellologdriver", 0)
	if err != nil {
		panic(err)
	}
}
