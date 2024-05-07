package rest

import (
	"fmt"
	"net/http"
)

type Mux struct {
	mux *http.ServeMux
}

func NewMux() *Mux {

	linkHandler := NewLinkHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", linkHandler.Link)
	return &Mux{
		mux: mux,
	}
}

func (m *Mux) ListenAndServe(addr string) error {
	if m.mux == nil {
		return fmt.Errorf("mux is not initialized")
	}
	return http.ListenAndServe(addr, m.mux)
}
