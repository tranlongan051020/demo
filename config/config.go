package config

import "time"

var (
	Server = struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		StaticDir    string
	}{
		Port:         "8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		StaticDir:    "./public",
	}
)
