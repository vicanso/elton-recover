# elton-recover

The middleware has been archived, please use the middleware of [elton](https://github.com/vicanso/elton).

[![Build Status](https://img.shields.io/travis/vicanso/elton-recover.svg?label=linux+build)](https://travis-ci.org/vicanso/elton-recover)


Recover middleware for elton, it can get panic error to avoid application crash.

```go
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
```
