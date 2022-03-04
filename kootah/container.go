package kootah

import (
	"kootah/definitions"
	"kootah/repositories"
	"log"
)

func New() definitions.Kootah {
	uidRep, err := repositories.NewUIDTeris()
	if err != nil {
		log.Fatalf("error while creating new uid repository: %v\n", err)
	}
	urlRep := repositories.NewURLInMemoryRepository()
	app := NewKootah(urlRep, uidRep)
	return app
}
