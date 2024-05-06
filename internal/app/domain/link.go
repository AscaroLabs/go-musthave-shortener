package domain

type ShortRequest struct {
	OriginalURL string
}

type ShortResponse struct {
	ShortURL string
}

type GetOriginalRequest struct {
	ID string
}

type GetOriginalResponse struct {
	OriginalURL string
}
