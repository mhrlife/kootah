package repositories

import (
	"github.com/teris-io/shortid"
	"kootah/definitions"
	"kootah/domain"
)

type UIDTeris struct {
	sid *shortid.Shortid
}

func NewUIDTeris() (definitions.UIDRepository, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return nil, err
	}
	return &UIDTeris{
		sid: sid,
	}, nil
}

func (U *UIDTeris) GetUID() (domain.UID, error) {
	uid, err := U.sid.Generate()
	if err != nil {
		return "", err
	}
	return domain.UID(uid), nil
}
