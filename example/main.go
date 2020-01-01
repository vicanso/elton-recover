package main

import (
	"errors"

	"github.com/vicanso/elton"

	recover "github.com/vicanso/elton-recover"
)

func main() {
	e := elton.New()

	e.Use(recover.New())

	e.GET("/", func(c *elton.Context) (err error) {
		panic(errors.New("abcd"))
	})

	err := e.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}
