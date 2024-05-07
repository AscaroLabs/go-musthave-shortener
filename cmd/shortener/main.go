package main

import (
	"flag"

	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/api/rest"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
)

func main() {
	flag.Parse()
	linkHandler := rest.NewLinkHandler()
	r := rest.NewRouter(
		linkHandler,
	)
	if err := r.Router().Run(*config.Addr); err != nil {
		panic(err)
	}
}
