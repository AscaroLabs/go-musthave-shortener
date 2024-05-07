package memory

import "sync"

type linkStore struct {
	mu          *sync.RWMutex
	mIDOriginal map[string]string
}

func NewLinkStore() *linkStore {
	return &linkStore{
		mu:          &sync.RWMutex{},
		mIDOriginal: make(map[string]string),
	}
}

func (s *linkStore) Save(id, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mIDOriginal[id] = originalURL
	return nil
}

func (s *linkStore) Get(id string) (*string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if val, ok := s.mIDOriginal[id]; ok {
		return &val, nil
	}
	return nil, nil
}
