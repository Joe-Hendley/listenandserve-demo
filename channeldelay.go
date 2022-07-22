package main

import (
	"net"
	"net/http"
)

func channelDelay(server *http.Server) {
	done := make(chan bool)
	defer close(done)

	go func() {
		l, err := net.Listen("tcp", port)
		if err != nil {
			panic(err)
		}

		done <- true

		err = server.Serve(l)
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-done
}
