package main

import (
	"kootah/delivery/rest"
	"kootah/kootah"
)

func main() {
	k := kootah.New()
	r := rest.NewRest(k)
	r.Listen()
}
