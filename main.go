package main

import (
	"net/http"
	"proxy/goproxy"
	"time"

	"log"
)

func main() {
	log.Print("服務起來了")
	proxy := goproxy.New()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      proxy,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
