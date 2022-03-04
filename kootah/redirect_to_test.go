package kootah

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"kootah/definitions"
	"kootah/domain"
	"kootah/mocks"
	"testing"
)

func TestKootah_GetRedirectTo(t *testing.T) {
	urlMock := &mocks.URLRepository{}
	ctx := context.Background()
	k := NewKootah(urlMock, nil)

	uid := domain.UID("uid")
	url := &domain.URL{
		ID:           uid,
		RedirectTo:   "https://google.com",
		VisitedCount: 12,
	}
	urlMock.On("Get", mock.Anything, uid).Return(url, nil)
	urlMock.On("Save", mock.Anything, mock.MatchedBy(func(url domain.URL) bool {
		return url.VisitedCount == 13
	})).Return(nil)

	redirectTo, err := k.GetRedirectTo(ctx, uid)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if redirectTo != url.RedirectTo {
		t.Fatalf("unexpected redirect to %s", redirectTo)
	}

	urlMock.AssertExpectations(t)
}

func TestKootah_GetRedirectToError(t *testing.T) {
	urlMock := &mocks.URLRepository{}
	ctx := context.Background()
	k := NewKootah(urlMock, nil)
	uid := domain.UID("uid")

	urlMock.On("Get", mock.Anything, uid).Return(nil, definitions.ErrEntityNotFound)

	_, err := k.GetRedirectTo(ctx, uid)
	if !errors.Is(err, definitions.ErrEntityNotFound) {
		t.Fatalf("unexpected error %v", err)
	}
	urlMock.AssertExpectations(t)
}
