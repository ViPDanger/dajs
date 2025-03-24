package main

import (
	"DAJ/Server/pkg/config"
	"DAJ/Server/pkg/logger"
)

func main() {
// 
	cfg:= config.NewConfig()
	_ = cfg.Set("User", "User")
	logger.Println(cfg.String("User"))
		// set value
		_ = cfg.Set("User", "User2")
		name := cfg.String("User")
		logger.Printf("- set string\n val: %v\n", name)
		logger.Println(cfg.String("User"))
}
