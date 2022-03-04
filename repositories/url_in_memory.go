package repositories

import (
	"context"
	"errors"
	"kootah/definitions"
	"kootah/domain"
	"sync"
)

var (
	ErrConvertionError = errors.New("couldn't convert map value to data")
)

type urlInMemoryRepository struct {
	data *sync.Map
}

func NewURLInMemoryRepository() definitions.URLRepository {
	return &urlInMemoryRepository{
		data: new(sync.Map),
	}
}

func (u *urlInMemoryRepository) Save(ctx context.Context, url domain.URL) error {
	u.data.Store(url.ID, &url)
	return nil
}

func (u *urlInMemoryRepository) Get(ctx context.Context, uid domain.UID) (*domain.URL, error) {
	urlInterface, exists := u.data.Load(uid)
	if !exists {
		return nil, definitions.ErrEntityNotFound
	}

	url, ok := urlInterface.(*domain.URL)
	if !ok {
		return nil, ErrConvertionError
	}
	return url, nil
}
