package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gokrazy/rsync/rsyncd"
)

func main() {
	// simulate user (or process supervisor) asking us to stop soon after starting
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	listener, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		log.Fatal(err)
	}

	rsyncServer, err := rsyncd.NewServer([]rsyncd.Module{
		{
			Name: "test",
			Path: "/test_upload",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := rsyncServer.Serve(ctx, listener); err != nil {
		log.Fatal(err)
	}

	log.Println("gracefully exiting")
}
