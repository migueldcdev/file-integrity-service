package main

import (
	"flag"
	"github.com/migueldcdev/file-integrity-service/internal/monitor"
)

func main() {
	path := flag.String("path", "", "dir path to watch")
	flag.Parse()
	monitor.WatchDirectory(*path)

}
