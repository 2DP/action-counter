package main

import (
	"github.com/2DP/action-counter/server"
	"github.com/2DP/action-counter/config"
)

func main() {
	config := &config.Config{
		RedisAddr:"localhost:6379",
		RedisPassword:""}	
	
	server := &server.Server{}
	server.Initialize(config)
	
	server.Run(":8080")
}

