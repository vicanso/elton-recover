# elton-recover

[![Build Status](https://img.shields.io/travis/vicanso/elton-recover.svg?label=linux+build)](https://travis-ci.org/vicanso/elton-recover)


Recover middleware for elton, it can get panic error to avoid application crash.

```go
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

	d.ListenAndServe(":3000")
}

```