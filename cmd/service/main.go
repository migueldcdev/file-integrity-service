package main

import (
	"github.com/migueldcdev/file-integrity-service/internal/monitor"
)

func main() {
	path := "/home/miguel/Dev/go/file-integrity-service/watchdir"
	monitor.WatchDirectory(path)

}
