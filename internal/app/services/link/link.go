package link

import (
	"fmt"

	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/store/memory"
	"github.com/AscaroLabs/go-musthave-shortener/internal/app/domain"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
	"github.com/AscaroLabs/go-musthave-shortener/internal/utils"
)

type Store interface {
	Save(id, originalURL string) error
	Get(id string) (*string, error)
}

type LinkService struct {
	store Store
}

func NewLinkService() *LinkService {
	return &LinkService{
		store: memory.NewLinkStore(),
	}
}

func (ls *LinkService) Short(req domain.ShortRequest) (*domain.ShortResponse, error) {
	id := generateId()
	if err := ls.store.Save(id, req.OriginalURL); err != nil {
		return nil, fmt.Errorf("failed to save link: %w", err)
	}
	return &domain.ShortResponse{
		ShortURL: urlById(id),
	}, nil
}

func (ls *LinkService) GetOriginal(req domain.GetOriginalRequest) (*domain.GetOriginalResponse, error) {
	originalUrl, err := ls.store.Get(req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get link(id=%s): %w", req.ID, err)
	}
	if originalUrl != nil {
		return &domain.GetOriginalResponse{
			OriginalURL: *originalUrl,
		}, nil
	}
	return nil, nil
}

func urlById(id string) string {
	return fmt.Sprintf("%s://%s%s/%s", config.NET_PROTOCOL, config.HTTP_HOST, config.HTTP_PORT, id)
}
func generateId() string {
	return utils.RandomString(config.ID_LENGTH)
}
