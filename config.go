package fdfs_client

import (
	"strconv"
)

type config struct {
	trackerAddr []string
	maxConns    int
}

func newConfigByYaml(track_server []string,maxcon string) (*config, error)  {
	cfg := &config{}
	cfg.maxConns, _ = strconv.Atoi(maxcon)
	cfg.trackerAddr = track_server

	return cfg, nil
}