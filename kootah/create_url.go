package kootah

import (
	"context"
	"kootah/domain"
)

func (k *Kootah) CreateURL(ctx context.Context, redirectTo string) (domain.UID, error) {
	// generates a new Unique ID
	uid, err := k.uids.GetUID()
	if err != nil {
		return "", err
	}
	// creates URL model
	url := domain.URL{
		ID:         uid,
		RedirectTo: redirectTo,
	}
	// saves created url in repository
	if err := k.urls.Save(ctx, url); err != nil {
		return "", err
	}
	return uid, nil

}
