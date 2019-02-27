# cod-recover

[![Build Status](https://img.shields.io/travis/vicanso/cod-recover.svg?label=linux+build)](https://travis-ci.org/vicanso/cod-recover)


Recover middleware for cod, it can get panic error to avoid application crash.

```go
package main

import (
	"github.com/vicanso/cod"

	recover "github.com/vicanso/cod-recover"
)

func main() {
	d := cod.New()

	d.Use(recover.New())

	d.GET("/", func(c *cod.Context) (err error) {
		panic("abcd")
	})

	d.ListenAndServe(":7001")
}

```