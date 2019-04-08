package recover

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vicanso/cod"
)

func TestRecover(t *testing.T) {
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
	if resp.Code != http.StatusInternalServerError ||
		resp.Body.String() != "category=cod-recover, message=abc" ||
		!ctx.Committed ||
		!catchError {
		t.Fatalf("recover fail")
	}
	for _, key := range keys {
		if resp.Header().Get(key) != "" {
			t.Fatalf("reset response header fail")
		}
	}
}
