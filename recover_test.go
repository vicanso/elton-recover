package recover

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/cod"
)

func TestRecover(t *testing.T) {
	assert := assert.New(t)
	var ctx *cod.Context
	d := cod.New()
	d.Use(New())
	d.GET("/", func(c *cod.Context) error {
		ctx = c
		panic("abc")
	})
	req := httptest.NewRequest("GET", "https://aslant.site/", nil)
	resp := httptest.NewRecorder()
	keys := []string{
		cod.HeaderETag,
		cod.HeaderLastModified,
		cod.HeaderContentEncoding,
		cod.HeaderContentLength,
	}
	for _, key := range keys {
		resp.Header().Set(key, "a")
	}

	catchError := false
	d.OnError(func(_ *cod.Context, _ error) {
		catchError = true
	})

	d.ServeHTTP(resp, req)
	assert.Equal(resp.Code, http.StatusInternalServerError)
	assert.Equal(resp.Body.String(), "category=cod-recover, message=abc")
	assert.True(ctx.Committed)
	assert.True(catchError)
	for _, key := range keys {
		assert.Empty(ctx.GetHeader(key), "header should be reseted")
	}
}

// https://stackoverflow.com/questions/50120427/fail-unit-tests-if-coverage-is-below-certain-percentage
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.9 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}
