package main

import (
	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/api/rest"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
)

func main() {
	mux := rest.NewMux()
	if err := mux.ListenAndServe(config.HTTP_PORT); err != nil {
		panic(err)
	}
}
