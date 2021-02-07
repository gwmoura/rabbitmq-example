package main

import (
	"g.com/logger"
)

func startLogger(program string) {
	if program == "console" {
		logger.StartConsole()
	} else if program == "writer" {
		logger.StartWriter()
	} else {
		logger.StartConsole()
	}
}
