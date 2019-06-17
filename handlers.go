package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/docker/docker/daemon/logger"
	"github.com/pkg/errors"
)

//StartLoggingRequest represents the request object we get on a call to //LogDriver.StartLogging
type StartLoggingRequest struct {
	File string
	Info logger.Info
}

//StopLoggingRequest represents the request object we get on a call to //LogDriver.StopLogging
type StopLoggingRequest struct {
	File string
}

//This gets called when a container starts that requests the log driver
func startLoggingHandler() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var startReq StartLoggingRequest
		err := json.NewDecoder(r.Body).Decode(&startReq)
		if err != nil {
			http.Error(w, errors.Wrap(err, "error decoding json request").Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(os.Stderr, "Got start request object from container %#v\n", startReq.Info.ContainerName)
		fmt.Fprintf(os.Stderr, "Got a container with the following labels: %#v\n", startReq.Info.ContainerLabels)
		fmt.Fprintf(os.Stderr, "Got a container with the following log opts: %#v\n", startReq.Info.Config)

		cfg, err := handleConfig(startReq.Info.Config)
		if err != nil {
			http.Error(w, errors.Wrap(err, "error creating plugin config").Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(os.Stderr, "Created config object %#v\n", cfg)
		logHandler, err := newLogger(startReq.File)
		if err != nil {
			http.Error(w, errors.Wrap(err, "error creating logger").Error(), http.StatusBadRequest)
			return
		}
		go logHandler.consumeLogs()

		respondOK(w)
	} //end func
}

//This gets called when a container using the log driver stops
func stopLoggingHandler() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var stopReq StopLoggingRequest
		err := json.NewDecoder(r.Body).Decode(&stopReq)
		if err != nil {
			http.Error(w, errors.Wrap(err, "error decoding json request").Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(os.Stderr, " Got stop request object %#v\n", stopReq)
		respondOK(w)
	} // end func
}

//For the start/stop handler, the daemon expects back an error object. If the body is empty, then all is well.
func respondOK(w http.ResponseWriter) {
	res := struct {
		Err string
	}{
		"",
	}

	json.NewEncoder(w).Encode(&res)
}
