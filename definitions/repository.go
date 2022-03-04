package definitions

import (
	"context"
	"kootah/domain"
)

type UIDRepository interface {
	GetUID() (domain.UID, error)
}

type URLRepository interface {
	Save(ctx context.Context, url domain.URL) error
	Get(ctx context.Context, uid domain.UID) (*domain.URL, error)
}
