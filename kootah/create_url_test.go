package kootah

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"kootah/domain"
	"kootah/mocks"
	"testing"
)

func TestKootah_CreateURL(t *testing.T) {
	urlMock := &mocks.URLRepository{}
	uidMock := &mocks.UIDRepository{}
	k := NewKootah(urlMock, uidMock)

	uidMock.On("GetUID").Return(domain.UID("uid"), nil)
	urlMock.On("Save", mock.Anything, mock.MatchedBy(func(url domain.URL) bool {
		if url.ID != "uid" {
			return false
		}
		return true
	})).Return(nil)

	ctx := context.Background()
	uid, err := k.CreateURL(ctx, "https://google.com")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if uid != "uid" {
		t.Fatalf("wrong uid, expected uid, got %s", uid)
	}

	uidMock.AssertExpectations(t)
	urlMock.AssertExpectations(t)
}

func TestKootah_CreateURLError(t *testing.T) {
	uidMock := &mocks.UIDRepository{}
	k := NewKootah(nil, uidMock)

	ctx := context.Background()
	returnError := errors.New("someError")

	uidMock.On("GetUID").Return(domain.UID(""), returnError)
	_, err := k.CreateURL(ctx, "https://google.com")

	if !errors.Is(err, returnError) {
		t.Fatalf("unexpected error %v", err)
	}
	uidMock.AssertExpectations(t)
}
