package main

import (
	"net"
	"net/http"
	"time"
)

func channelDelayTimeout(server *http.Server) {
	done := make(chan struct{})
	defer close(done)

	go func() {
		l, err := net.Listen("tcp", port)
		if err != nil {
			panic(err)
		}

		done <- struct{}{}

		err = server.Serve(l)
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	select {
	case <-done:
		return
	case <-time.After(1 * time.Second):
		panic("server timed out")
	}
}
