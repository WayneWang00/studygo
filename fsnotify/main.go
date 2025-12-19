package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					log.Println("created file:", event.Name)
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
				if event.Has(fsnotify.Remove) {
					log.Println("removed file:", event.Name)
				}
				if event.Has(fsnotify.Rename) {
					log.Println("renamed file:", event.Name)
				}
				if event.Has(fsnotify.Chmod) {
					log.Println("chmod file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
			log.Println("watch list:", watcher.WatchList())
		}
	}()

	// Add a path.
	err = watcher.Add("D:\\Code\\Go\\personal")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
