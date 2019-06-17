# Hellologdriver: a simple, hello-world style example of a docker logging plugin


This repo represents the most basic things you will need in order to have a working [docker logging plugin](https://docs.docker.com/engine/extend/plugins_logging/). It's heavily instrumented with log lines, and will output any logs it receives to stderr. It also demonstrates how to use any config passed via `--log-opts`

## Build and install

To build and install, just run `make all`


## Running

`docker run --log-driver=ossifrage/hellologdriver:0.0.1 --log-opt hostname=localhost -it debian:jessie /bin/bash`


## How it works

Logging plugins work by starting up an HTTP server that reads over a unix socket. When a container starts up that requests the logging plugin, a request is sent to `/LogDriver.StartLogging` with the name of the log handler and a struct containing the config of the container, including labels and other metadata. The actual log reading requires the file handle to be passed to a new routine which uses protocol buffers to read from the log handler. When the container stops, a request is sent to `/LogDriver.StopLogging`.