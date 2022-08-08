package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	addr     = "http://localhost"
	port     = ":9000"
	attempts = 1000
)

func main() {
	http.HandleFunc("/", handleOK)

	successes := 0
	now := time.Now()

	for i := 0; i < attempts; i++ {
		server := &http.Server{
			Addr: port,
		}

		successes += tryPing(
			server,
			noDelay,
		)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	then := time.Now()
	diff := then.Sub(now)

	fmt.Printf("%d successes out of %d attempts in %v\n", successes, attempts, diff)
}

func tryPing(server *http.Server, startServer func(*http.Server)) int {

	startServer(server)

	resp, err := http.Get(addr + port)
	if err != nil {
		return 0
	}

	defer resp.Body.Close()
	return 1
}
