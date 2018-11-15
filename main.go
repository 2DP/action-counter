package main

import (
	"github.com/2DP/action-counter/server"
	"github.com/2DP/action-counter/config"
)

func main() {
	config := &config.Config{}	
	server := &server.Server{}
	server.Initialize(config)
	server.Run(":8080")
}

