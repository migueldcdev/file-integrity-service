package main

import (
	"flag"
	"fmt"
	"github.com/migueldcdev/file-integrity-service/internal/db"
	"github.com/migueldcdev/file-integrity-service/internal/hash"
	"github.com/migueldcdev/file-integrity-service/internal/monitor"
	"os"
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

	hash.WalkDirAndHashFiles(*path)
	monitor.WatchDirectory(*path)

}
