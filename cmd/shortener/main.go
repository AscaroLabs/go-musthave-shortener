package main

import (
	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/api/rest"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
)

func main() {
	if !config.Initialized() {
		config.Init()
	}
	linkHandler := rest.NewLinkHandler()
	r := rest.NewRouter(
		linkHandler,
	)
	if err := r.Router().Run(config.Config.Addr); err != nil {
		panic(err)
	}
}
