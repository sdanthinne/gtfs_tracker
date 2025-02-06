package cmd

import (
    "log"
    "sdanthinne/gtfs_tracker/internal/server"
)

func validate_args(args [] string) bool {
	if len(args) > 1 {
		return false
	}
	return true
}

func Run(args [] string) int {
	if !validate_args(args) {
		return -1
	}
	log.Printf("Initializing...")
	server := Create_server()

	Start_server(server)
	return 0
}

