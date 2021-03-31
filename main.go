package main

import (
	"gitlab.com/F1ukez/sample-go/router"
)

func main() {
	r := router.New()
	router.Mount(r)
	r.Run()
}
