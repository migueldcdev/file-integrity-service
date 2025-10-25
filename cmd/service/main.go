package main

import (
	"flag"
	"fmt"
	"github.com/migueldcdev/file-integrity-service/internal/api"
	"github.com/migueldcdev/file-integrity-service/internal/db"
	"github.com/migueldcdev/file-integrity-service/internal/monitor"
	"github.com/migueldcdev/file-integrity-service/internal/service"
	"os"
	"os/signal"
	"syscall"
)

const DB_PATH string = "../file-integrity-service/data/sqlite.db"

func main() {
	if err := db.ConnectToDB(DB_PATH); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while connecting to DB: %v\n", err)
		os.Exit(1)
	}

	if err := db.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while initializing DB: %v\n", err)
		os.Exit(1)
	}

	path := flag.String("path", "", "dir path to watch")
	flag.Parse()

	go api.RunAPIServer()
	go monitor.WatchDirectory(*path)
	service.WalkDirAndHashFiles(*path)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nService terminated")
}
