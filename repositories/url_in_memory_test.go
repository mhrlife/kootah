package repositories

import (
	"context"
	"kootah/domain"
	"testing"
)

func TestUrlInMemoryRepository_Get(t *testing.T) {
	r := NewURLInMemoryRepository()
	url := domain.URL{
		ID:         domain.UID("123"),
		RedirectTo: "https://google.com",
	}
	_ = r.Save(context.Background(), url)
	fetchedUrl, err := r.Get(context.Background(), url.ID)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if fetchedUrl.RedirectTo != url.RedirectTo {
		t.Fatalf("unexpected url %v", fetchedUrl)
	}
}
