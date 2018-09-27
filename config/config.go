package config

import (
	"clip/cmd"
	"clip/config"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	config.Config
	ClipService Service
}

type Service struct {
	Address string
}

func init() {
	// Init CLI commands
	cmd.Root().Use = "bin/stream --config <Config path>"
	cmd.Root().Short = "stream - Provide API for stream"
	cmd.Root().Long = "stream"

	cmd.SetRunFunc(load)
}

func load() {
	once.Do(func() {
		mConfig := config.Load()
		if err := cmd.GetViper().Unmarshal(&conf); err != nil {
			panic(err)
		}
		conf.Config = mConfig
	})
}

func Load() {
	load()
}

func Get() Config {
	load()
	return conf
}
