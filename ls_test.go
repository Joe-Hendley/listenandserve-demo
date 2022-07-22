package main

import (
	"context"
	"net/http"
	"testing"
)

func BenchmarkNoDelay(b *testing.B) {
	successes := 0
	for i := 0; i < b.N; i++ {
		server := &http.Server{
			Addr: port,
		}
		successes = successes + tryPing(server, noDelay)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	b.Logf("%d successes out of %d attempts\n", successes, b.N)
}

func BenchmarkShortSleepDelay(b *testing.B) {
	successes := 0
	for i := 0; i < b.N; i++ {
		server := &http.Server{
			Addr: port,
		}
		successes = successes + tryPing(server, sleepDelay)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	b.Logf("%d successes out of %d attempts\n", successes, b.N)
}
func BenchmarkLongSleepDelay(b *testing.B) {
	successes := 0
	for i := 0; i < b.N; i++ {
		server := &http.Server{
			Addr: port,
		}
		successes = successes + tryPing(server, longSleepDelay)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	b.Logf("%d successes out of %d attempts\n", successes, b.N)
}
func BenchmarkChannelDelay(b *testing.B) {
	successes := 0
	for i := 0; i < b.N; i++ {
		server := &http.Server{
			Addr: port,
		}
		successes = successes + tryPing(server, channelDelay)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	b.Logf("%d successes out of %d attempts\n", successes, b.N)
}
func BenchmarkChannelTimeoutDelay(b *testing.B) {
	successes := 0
	for i := 0; i < b.N; i++ {
		server := &http.Server{
			Addr: port,
		}
		successes = successes + tryPing(server, channelDelayTimeout)

		err := server.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
	b.Logf("%d successes out of %d attempts\n", successes, b.N)
}
