package monitor

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

func WatchDirectory(path string) {
	w, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Printf("An error ocurred while creating a new watcher %s", err)
		os.Exit(1)
	}

	w.Add(path)

	go watchLoop(w)

	defer w.Close()
	fmt.Println("Service is running, press <Ctrl+c> to exit")
	<-make(chan struct{})
}

func watchLoop(w *fsnotify.Watcher) {
	i := 0
	for {
		select {

		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			fmt.Printf("ERROR: %s\n", err)

		case e, ok := <-w.Events:
			if !ok {
				return
			}
			i++
			fmt.Printf("%3d %s\n", i, e)
		}
	}
}
