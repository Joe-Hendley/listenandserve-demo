package main

import (
	"net/http"
)

func noDelay(server *http.Server) {
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			return
		}
	}()
}
