package definitions

import (
	"context"
	"kootah/domain"
)

type Kootah interface {
	CreateURL(ctx context.Context, redirectTo string) (domain.UID, error)
	GetRedirectTo(ctx context.Context, urlID domain.UID) (string, error)
}
