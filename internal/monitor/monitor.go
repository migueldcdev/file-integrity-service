package monitor

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io/fs"
	"os"
	"path/filepath"
)

func WatchDirectory(path string) {
	w, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Printf("An error ocurred while creating a new watcher %s", err)
		os.Exit(1)
	}

	if err = addRecursive(w, path); err != nil {
		fmt.Printf("Error adding path to watcher: %v", err)
		os.Exit(1)
	}

	defer w.Close()

	fmt.Println("Service is running, press <Ctrl+c> to exit")

	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}
			fmt.Printf("EVENT: %s %s\n", event.Op, event.Name)

		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func addRecursive(w *fsnotify.Watcher, path string) error {
	return filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if err := w.Add(path); err != nil {
				fmt.Printf("Error adding path %s: %v\n", path, err)
			} else {
				fmt.Printf("Watching: %s\n", path)
			}
		}
		return nil
	})
}
