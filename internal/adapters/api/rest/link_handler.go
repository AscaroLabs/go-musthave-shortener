package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/AscaroLabs/go-musthave-shortener/internal/app/domain"
	"github.com/AscaroLabs/go-musthave-shortener/internal/app/services/link"
)

type linkHandler struct {
	linkService *link.LinkService
}

func NewLinkHandler() *linkHandler {
	return &linkHandler{
		linkService: link.NewLinkService(),
	}
}

func (h *linkHandler) Link(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		h.Short(w, req)
	case http.MethodGet:
		h.RedirectOriginal(w, req)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *linkHandler) Short(w http.ResponseWriter, req *http.Request) {
	originalURL, err := getUrlFromBody(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.linkService.Short(domain.ShortRequest{
		OriginalURL: originalURL,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if res != nil {
		writeShortResponse(*res, w)
	}
}

func (h *linkHandler) RedirectOriginal(w http.ResponseWriter, req *http.Request) {
	pathParts := strings.Split(req.URL.Path, "/")
	if len(pathParts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := pathParts[1]
	res, err := h.linkService.GetOriginal(domain.GetOriginalRequest{
		ID: id,
	})
	if err != nil || res == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	writeRedirectOriginalResponse(*res, w)
}

func getUrlFromBody(body io.Reader) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	urlText := string(bodyBytes)
	// Валидация
	if _, err := url.Parse(urlText); err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	return urlText, nil
}

func writeShortResponse(res domain.ShortResponse, w http.ResponseWriter) {
	w.Write([]byte(res.ShortURL))
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
}

func writeRedirectOriginalResponse(res domain.GetOriginalResponse, w http.ResponseWriter) {
	w.Header().Add("Location", res.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
