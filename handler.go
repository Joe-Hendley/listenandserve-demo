package main

import (
	"net/http"
)

func handleOK(w http.ResponseWriter, req *http.Request) {
	resp := http.Response{
		StatusCode: http.StatusOK,
	}
	_ = resp.Write(w)
}
