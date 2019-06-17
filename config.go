package main

import "fmt"

//HelloCfg is the docker plugin config object, specified by the user with --log-opt
type HelloCfg struct {
	Hostname string
}

//Check the config object that came from docker
func handleConfig(dockerCfg map[string]string) (HelloCfg, error) {
	var cfg HelloCfg
	esHostname, ok := dockerCfg["hostname"]
	if !ok {
		return cfg, fmt.Errorf("container must have an hostname log option")
	}
	cfg.Hostname = esHostname

	return cfg, nil
}
