package main

import (
	"github.com/vicanso/elton"

	recover "github.com/vicanso/elton-recover"
)

func main() {
	d := elton.New()

	d.Use(recover.New())

	d.GET("/", func(c *elton.Context) (err error) {
		panic("abcd")
	})

	err := d.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}
