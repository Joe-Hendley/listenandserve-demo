package main

import (
	"net/http"
	"time"
)

func sleepDelay(server *http.Server) {

	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			return
		}
	}()

	time.Sleep(1 * time.Nanosecond)
}

func longSleepDelay(server *http.Server) {

	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			return
		}
	}()

	time.Sleep(1 * time.Millisecond)
}
