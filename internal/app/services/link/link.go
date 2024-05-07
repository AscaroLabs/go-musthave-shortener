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
	id := generateID()
	if err := ls.store.Save(id, req.OriginalURL); err != nil {
		return nil, fmt.Errorf("failed to save link: %w", err)
	}
	return &domain.ShortResponse{
		ShortURL: urlByID(id),
	}, nil
}

func (ls *LinkService) GetOriginal(req domain.GetOriginalRequest) (*domain.GetOriginalResponse, error) {
	originalURL, err := ls.store.Get(req.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get link(id=%s): %w", req.ID, err)
	}
	if originalURL != nil {
		return &domain.GetOriginalResponse{
			OriginalURL: *originalURL,
		}, nil
	}
	return nil, nil
}

func urlByID(id string) string {
	return *config.Base + "/" + id
}
func generateID() string {
	return utils.RandomString(config.IDLength)
}
