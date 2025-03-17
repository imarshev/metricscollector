package main

import (
	"github.com/imarshev/metricscollector/internal/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", server.UpdateHandler)
	http.ListenAndServe("localhost:8080", mux)
}
