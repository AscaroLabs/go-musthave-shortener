package main

import (
	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/api/rest"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
)

func main() {
	linkHandler := rest.NewLinkHandler()
	r := rest.NewRouter(
		linkHandler,
	)
	if err := r.Router().Run(config.HTTPHost + config.HTTPPort); err != nil {
		panic(err)
	}
}
