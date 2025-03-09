package main

import (
	"log"
	"os"
	"sdanthinne/gtfs_tracker/cmd"
)

func main() {
	log.SetFlags(0)
	cmd.Run(os.Args)
}
