package kootah

import (
	"kootah/definitions"
)

type Kootah struct {
	urls definitions.URLRepository
	uids definitions.UIDRepository
}

func NewKootah(urls definitions.URLRepository, uids definitions.UIDRepository) definitions.Kootah {
	return &Kootah{
		uids: uids,
		urls: urls,
	}
}
