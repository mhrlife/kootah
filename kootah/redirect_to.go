package kootah

import (
	"context"
	"github.com/sirupsen/logrus"
	"kootah/domain"
)

func (k *Kootah) GetRedirectTo(ctx context.Context, urlID domain.UID) (string, error) {
	url, err := k.urls.Get(ctx, urlID)
	if err != nil {
		return "", err
	}
	url.VisitedCount++
	err = k.urls.Save(ctx, *url)
	if err != nil {
		logrus.WithError(err).Error("error while incrementing url visit")
	}
	return url.RedirectTo, nil
}
