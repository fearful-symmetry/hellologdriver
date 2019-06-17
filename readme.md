# Hellologdriver: a simple, hello-world style example of a docker logging plugin


This repo represents the most basic things you will need in order to have a working [docker logging plugin](https://docs.docker.com/engine/extend/plugins_logging/). It's heavily instrumented with log lines, and will output any logs it receives to stderr. It also demonstrates how to use any config passed via `--log-opts`

## Build and install

To build and install, just run `make all`


## Running

`docker run --log-driver=ossifrage/hellologdriver:0.0.1 --log-opt hostname=localhost -it debian:jessie /bin/bash`


## How it works

Logging plugins work by starting up an HTTP server that reads over a unix socket. When a container starts up that requests the logging plugin, a request is sent to `/LogDriver.StartLogging` with the name of the log handler and a struct containing the config of the container, including labels and other metadata. The actual log reading requires the file handle to be passed to a new routine which uses protocol buffers to read from the log handler. When the container stops, a request is sent to `/LogDriver.StopLogging`.



## if you're on OSX, it's nearly impossible to debug docker plugins and the docs are wrong 60% of the time.

First, you need to shell into the VM that `runc` lives in. To do this, you need to find the tty for the VM. On newer versions of Docker For Mac, it's at: `~/Library/Containers/com.docker.docker/Data/vms/0/tty` On older versions, it's _somewhere_ else in `~/Library/Containers/com.docker.docker`. 


Once you find it, just run `screen ~/Library/Containers/com.docker.docker/Data/vms/0/tty`


The location of the logs AND the container base directory in the docker docs is wrong.


You can use this to find the list of plugins running on the host: `runc --root /containers/services/docker/rootfs/run/docker/plugins/runtime-root/plugins.moby/ list`

The logs are in `/var/log/docker`. If you want to make the logs useful, you need to find the ID of the plugin. Back on the darwin host, run `docker plugin inspect $name_of_plugin | grep Id` use the hash ID to grep through the logs: `grep 22bb02c1506677cd48cc1cfccc0847c1b602f48f735e51e4933001804f86e2e docker.*`